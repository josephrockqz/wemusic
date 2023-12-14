package services

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/josephrockqz/wemusic-golang/internal/status"
	"github.com/josephrockqz/wemusic-golang/pkg/app"
)

func SpotifyUserAuthorizationCallback(context *gin.Context) {
	appG := app.Gin{C: context}

	queryParams := context.Request.URL.Query()

	code, ok := queryParams["code"]
	if !ok {
		fmt.Println("Could not get Spotify user authorization code from request URL.")
		appG.Response(http.StatusServiceUnavailable, status.ERROR, map[string]interface{}{
			"message": "could not get Spotify user authorization code from request URL.",
		})
		return
	}

	accessToken, err := GetAccessToken(code[0])
	if err != nil {
		fmt.Println("Could not get Spotify access token.")
		appG.Response(http.StatusServiceUnavailable, status.ERROR, map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	fmt.Println("access token:", accessToken)
	appG.Response(http.StatusAccepted, status.SUCCESS, map[string]interface{}{
		"access_token": accessToken,
	})
}
