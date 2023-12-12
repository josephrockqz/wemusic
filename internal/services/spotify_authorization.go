package services

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/josephrockqz/wemusic-golang/internal/status"
	"github.com/josephrockqz/wemusic-golang/pkg/app"
	"github.com/josephrockqz/wemusic-golang/pkg/utils"
)

func SpotifyUserAuthorization(context *gin.Context) {
	appG := app.Gin{C: context}
	context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	context.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	context.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	clientId := os.Getenv("SPOTIFY_CLIENT_ID")
	if clientId == "" {
		fmt.Println("Could not get Spotify client id. Please set SPOTIFY_CLIENT_ID environment variable.")
		appG.Response(http.StatusServiceUnavailable, status.ERROR, map[string]interface{}{
			"message": "could not get Spotify client id",
		})
		return
	}

	state := utils.GenerateRandomString(16)
	url := "https://accounts.spotify.com/authorize?client_id=" + clientId + "&response_type=code&redirect_uri=http://localhost:8080/spotify-user-authorization-callback&scope=user-read-private%20user-read-email&state=" + state

	context.Redirect(http.StatusTemporaryRedirect, url)
}

func SpotifyUserAuthorizationCallback(context *gin.Context) {
	appG := app.Gin{C: context}
	context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	context.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	context.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	fmt.Println("Spotify user authorization callback")

	appG.Response(http.StatusOK, status.SUCCESS, map[string]interface{}{
		"message": "spotifyUserAuthorizationCallback",
	})
}
