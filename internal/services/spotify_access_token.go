package services

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

// TODO: refactor to use echo framework instead of gin
// https://developer.spotify.com/documentation/web-api/tutorials/code-flow
func GetAccessToken(context echo.Context, code string) (string, error) {
	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", nil)
	if err != nil {
		context.Logger().Error("could not create Spotify access token request:", err)
		return "", echo.NewHTTPError(http.StatusInternalServerError, "could not create Spotify access token request:", err)
	}

	clientId := os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")
	if clientId == "" || clientSecret == "" {
		context.Logger().Error("Could not get Spotify client id or client secret. Please set SPOTIFY_CLIENT_ID and SPOTIFY_CLIENT_SECRET environment variables.")
		return "", echo.NewHTTPError(http.StatusInternalServerError, "Could not get Spotify client id or client secret from environment")
	}

	query := req.URL.Query()
	query.Add("grant_type", "authorization_code")
	query.Add("code", code)
	query.Add("redirect_uri", "http://localhost:8080/spotify-user-authorization-callback")
	req.URL.RawQuery = query.Encode()

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	clientCredentials := clientId + ":" + clientSecret
	clientCredentialsEncoding := base64.StdEncoding.EncodeToString([]byte(clientCredentials))
	req.Header.Set("Authorization", "Basic "+clientCredentialsEncoding)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		context.Logger().Error("could not get response for Spotify access token request:", err)
		return "", echo.NewHTTPError(http.StatusInternalServerError, "could not get response for Spotify access token request:", err)
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		context.Logger().Error("could not read response body for Spotify access token:", err)
		return "", echo.NewHTTPError(http.StatusInternalServerError, "could not read response body for Spotify access token:", err)
	}

	// TODO: use structs when unmarshaling response

	// type AccessTokenResponseSuccessData struct {
	// 	AccessToken  string `json:"access_token"`
	// 	TokenType    string `json:"token_type"`
	// 	ExpiresIn    int    `json:"expires_in"`
	// 	RefreshToken string `json:"refresh_token"`
	// }

	// type AccessTokenResponseFailureData struct {
	// 	Error            string `json:"error"`
	// 	ErrorDescription string `json:"error_description"`
	// }

	var accessTokenResponseData map[string]interface{}
	err = json.Unmarshal(respBody, &accessTokenResponseData)
	if err != nil {
		context.Logger().Error("could not unmarshal response body for Spotify access token:", err)
		return "", echo.NewHTTPError(http.StatusInternalServerError, "could not unmarshal response body for Spotify access token:", err)
	}

	if error, ok := accessTokenResponseData["error"]; ok {
		context.Logger().Error("error in Spotify access token response body:", error, accessTokenResponseData["error_description"])
		return "", echo.NewHTTPError(http.StatusInternalServerError, "error in Spotify access token response body:", error, accessTokenResponseData["error_description"])
	}
	if accessToken, ok := accessTokenResponseData["access_token"]; !ok {
		context.Logger().Error("could not get access token Spotify access token response body:", err)
		return "", echo.NewHTTPError(http.StatusInternalServerError, "could not get access token Spotify access token response body:", err)
	} else {
		return accessToken.(string), nil
	}
}
