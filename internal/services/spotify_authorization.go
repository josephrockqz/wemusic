package services

import (
	"errors"
	"fmt"

	"github.com/labstack/echo/v4"
)

func SpotifyUserAuthorizationCallback(context echo.Context) error {
	queryParams := context.QueryParams()
	fmt.Println(queryParams)

	state, ok := queryParams["state"]
	if !ok {
		fmt.Println("Could not get state from request URL.")
		return errors.New("Could not get state from request URL.")
	}
	// TODO: compare state to stored state
	fmt.Println("state:", state)

	code, ok := queryParams["code"]
	fmt.Println("code:", code)
	if !ok {
		fmt.Println("Could not get Spotify user authorization code from request URL.")
		return errors.New("Could not get Spotify user authorization code from request URL.")
	}

	accessToken, err := GetAccessToken(code[0])
	if err != nil {
		fmt.Println("Could not get Spotify access token.")
		return errors.New("Could not get Spotify access token.")
	}

	// TODO: make Spotify Library API call w/access token
	// https://github.com/spotify/web-api-examples/blob/7c4872d343a6f29838c437cf163012947b4bffb9/authorization/authorization_code/app.js#L84
	// can either make call in back end or in browser

	fmt.Println("access token:", accessToken)
	return nil
}
