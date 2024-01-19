package services

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func SpotifyUserAuthorizationCallback(context echo.Context) error {
	queryParams := context.QueryParams()

	state, ok := queryParams["state"]
	if !ok {
		context.Logger().Error("Could not get state from request URL.")
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not get state from request URL.")
	}

	// TODO: compare state to stored state
	context.Logger().Info("state:", state)

	code, ok := queryParams["code"]
	if !ok {
		context.Logger().Error("Could not get Spotify user authorization code from request URL.")
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not get Spotify user authorization code from request URL.")
	}

	accessToken, err := GetAccessToken(context, code[0])
	if err != nil {
		context.Logger().Error("Could not get Spotify access token.")
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not get Spotify access token.")
	}

	// TODO: store access token for later use (as cookie?)
	context.Logger().Info("access token:", accessToken)

	return context.NoContent(http.StatusCreated)
}
