package services

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/josephrockqz/wemusic-golang/pkg/utils"
)

func SpotifyLogin(context *gin.Context) {
	redirectLocation := constructRedirectLocation()

	fmt.Println(redirectLocation)

	context.String(200, redirectLocation)
	// context.Redirect(http.StatusMovedPermanently, redirectLocation)
}

func constructRedirectLocation() string {
	clientId := os.Getenv("SPOTIFY_CLIENT_ID")
	state := utils.GenerateRandomString(16)

	url := "https://accounts.spotify.com/authorize"
	url += "?client_id=" + clientId
	url += "&response_type=code"
	url += "&redirect_uri=http://localhost:8080/spotify-user-authorization-callback"
	url += "&scope=user-read-private%20user-read-email"
	url += "&state=" + state

	return url
}
