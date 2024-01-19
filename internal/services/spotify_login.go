package services

import (
	"errors"
	"net/http"
	"os"

	"github.com/josephrockqz/wemusic-golang/pkg/utils"
	"github.com/labstack/echo/v4"
)

func SpotifyLogin(context echo.Context) error {
	redirectLocation, err := constructRedirectLocation(context)
	if err != nil {
		return errors.New("could not construct Spotify login redirect URI")
	}
	return context.Redirect(http.StatusPermanentRedirect, redirectLocation)
}

func constructRedirectLocation(context echo.Context) (string, error) {
	clientId := os.Getenv("SPOTIFY_CLIENT_ID")
	if clientId == "" {
		context.Logger().Error("Could not get Spotify client id. Please set SPOTIFY_CLIENT_ID environment variable.")
		return "", errors.New("Could not get Spotify client id from environment")
	}

	state := utils.GenerateRandomString(16)

	url := "https://accounts.spotify.com/authorize"
	url += "?client_id=" + clientId
	url += "&response_type=code"
	url += "&redirect_uri=http://localhost:8080/spotify-user-authorization-callback"
	url += "&scope=user-read-private%20user-read-email"
	url += "&state=" + state

	return url, nil
}
