package services

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SpotifyUserAuthorizationCallback(context echo.Context) error {
	queryParams := context.QueryParams()

	state, ok := queryParams["state"]
	if !ok {
		context.Logger().Error("Could not get state from request URL.")
		return errors.New("Could not get state from request URL.")
	}

	// TODO: compare state to stored state
	context.Logger().Print("state:", state)

	code, ok := queryParams["code"]
	if !ok {
		context.Logger().Error("Could not get Spotify user authorization code from request URL.")
		return errors.New("Could not get Spotify user authorization code from request URL.")
	}

	accessToken, err := GetAccessToken(context, code[0])
	if err != nil {
		context.Logger().Error("Could not get Spotify access token.")
		return errors.New("Could not get Spotify access token.")
	}

	err = GetLibrary(context, accessToken)
	if err != nil {
		context.Logger().Error("Get Spotify library function call failed", err)
	}

	// TODO: store access token for later use (as cookie?)
	context.Logger().Print("access token:", accessToken)

	return context.NoContent(http.StatusOK)
}
