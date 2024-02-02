package services

import (
	"net/http"

	"github.com/josephrockqz/wemusic-golang/pkg/utils"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func SpotifyLogin(context echo.Context) error {
	redirectLocation, err := constructRedirectLocation(context)
	if err != nil {
		return err
	}
	return context.Redirect(http.StatusPermanentRedirect, redirectLocation)
}

func constructRedirectLocation(context echo.Context) (string, error) {
	spotifyClientId, err := utils.GetEnvironmentVariableByName("SPOTIFY_CLIENT_ID")
	if spotifyClientId == "" {
		zap.L().Error("Could not get Spotify client id. Please set SPOTIFY_CLIENT_ID environment variable.")
		return "", echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid Spotify Client ID")
	}
	if err != nil {
		zap.L().Error("Error retrieving Spotify Client ID from config file")
		return "", echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid Spotify Client ID")
	}

	state := utils.GenerateRandomString(16)

	// cookie := new(http.Cookie)
	// cookie.Name = "spotify_authorize_state"
	// cookie.Value = state
	cookie := &http.Cookie{
		Name:  "spotify_authorize_state",
		Value: state,
	}
	context.SetCookie(cookie)

	url := "https://accounts.spotify.com/authorize"
	url += "?client_id=" + spotifyClientId
	url += "&response_type=code"
	url += "&redirect_uri=http://localhost:8080/spotify-user-authorization-callback"
	url += "&scope=user-read-private%20user-read-email"
	url += "&state=" + state

	return url, nil
}
