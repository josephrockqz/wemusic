package services

import (
	"errors"
	"fmt"

	"github.com/labstack/echo/v4"
)

func SpotifyUserAuthorizationCallback(context echo.Context) error {
	queryParams := context.QueryParams()

	state, ok := queryParams["state"]
	if !ok {
		fmt.Println("Could not get state from request URL.")
		return errors.New("Could not get state from request URL.")
	}
	// TODO: compare state to stored state
	fmt.Println("state:", state)

	code, ok := queryParams["code"]
	if !ok {
		fmt.Println("Could not get Spotify user authorization code from request URL.")
		return errors.New("Could not get Spotify user authorization code from request URL.")
	}

	accessToken, err := GetAccessToken(code[0])
	if err != nil {
		fmt.Println("Could not get Spotify access token.")
		return errors.New("Could not get Spotify access token.")
	}

	fmt.Println("access token:", accessToken)
	return nil
}
