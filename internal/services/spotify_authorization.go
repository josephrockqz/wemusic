package services

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func SpotifyUserAuthorizationCallback(context echo.Context) error {
	queryParams := context.QueryParams()

	state, ok := queryParams["state"]
	if !ok {
		zap.L().Error("Could not get state from request URL.")
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not get state from request URL.")
	}

	cookie, err := context.Cookie("spotify_authorize_state")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unable to retrieve state value from cookie")
	}
	if cookie.Value != state[0] {
		zap.L().Error("State value does not match state stored in cookie.")
		return echo.NewHTTPError(http.StatusInternalServerError, "State value does not match state stored in cookie.")
	}

	code, ok := queryParams["code"]
	if !ok {
		zap.L().Error("Could not get Spotify user authorization code from request URL.")
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not get Spotify user authorization code from request URL.")
	}

	accessToken, err := GetAccessToken(context, code[0])
	if err != nil {
		zap.L().Error("Could not get Spotify access token.")
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not get Spotify access token.")
	}

	// TODO: store access token for later use (as cookie?)
	zap.L().Info("access token: " + accessToken)

	return context.NoContent(http.StatusCreated)
}
