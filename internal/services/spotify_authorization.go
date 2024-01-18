package services

import (
	"errors"

	"github.com/labstack/echo/v4"
)

func SpotifyUserAuthorizationCallback(context echo.Context) error {
	context.Logger().Print("enter SpotifyUserAuthorizationCallback")

	queryParams := context.QueryParams()

	state, ok := queryParams["state"]
	if !ok {
		context.Logger().Print("Could not get state from request URL.")
		return errors.New("Could not get state from request URL.")
	}

	// TODO: compare state to stored state
	context.Logger().Print("state:", state)

	code, ok := queryParams["code"]
	if !ok {
		context.Logger().Print("Could not get Spotify user authorization code from request URL.")
		return errors.New("Could not get Spotify user authorization code from request URL.")
	}

	accessToken, err := GetAccessToken(context, code[0])
	if err != nil {
		context.Logger().Print("Could not get Spotify access token.")
		return errors.New("Could not get Spotify access token.")
	}

	err = GetLibrary(context, accessToken)
	if err != nil {
		context.Logger().Print("Get Spotify library function call failed", err)
	}

	// TODO: store access token for later use (as cookie?)
	context.Logger().Print("access token:", accessToken)

	return nil
}
