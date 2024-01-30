package services

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"

	"github.com/josephrockqz/wemusic-golang/internal/transport"
	"github.com/josephrockqz/wemusic-golang/pkg/utils"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// TODO: refactor to use echo framework instead of gin
// https://developer.spotify.com/documentation/web-api/tutorials/code-flow
func GetAccessToken(context echo.Context, code string) (string, error) {
	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", nil)
	if err != nil {
		zap.L().Error("could not create Spotify access token request: " + err.Error())
		return "", echo.NewHTTPError(http.StatusInternalServerError, "could not create Spotify access token request:", err)
	}

	clientId, err := utils.GetEnvironmentVariable("spotify_client_id")
	if err != nil {
		zap.L().Error("Error retrieving Spotify Client ID from config file")
		return "", echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid Spotify Client ID")
	}
	clientSecret, err := utils.GetEnvironmentVariable("spotify_client_secret")
	if err != nil {
		zap.L().Error("Error retrieving Spotify Client Secret from config file")
		return "", echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid Spotify Client ID")
	}
	if clientId == "" || clientSecret == "" {
		zap.L().Error("Could not get Spotify client id or client secret. Please set SPOTIFY_CLIENT_ID and SPOTIFY_CLIENT_SECRET environment variables.")
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

	client := transport.HttpClient()

	resp, err := client.Do(req)
	if err != nil {
		zap.L().Error("could not get response for Spotify access token request: " + err.Error())
		return "", echo.NewHTTPError(http.StatusInternalServerError, "could not get response for Spotify access token request:", err)
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		zap.L().Error("could not read response body for Spotify access token: " + err.Error())
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
		zap.L().Error("could not unmarshal response body for Spotify access token: " + err.Error())
		return "", echo.NewHTTPError(http.StatusInternalServerError, "could not unmarshal response body for Spotify access token:", err)
	}

	if _, ok := accessTokenResponseData["error"]; ok {
		zap.L().Error("error in Spotify access token response body")
		return "", echo.NewHTTPError(http.StatusInternalServerError, "error in Spotify access token response body:", accessTokenResponseData["error_description"])
	}
	if accessToken, ok := accessTokenResponseData["access_token"]; !ok {
		zap.L().Error("could not get access token Spotify access token response body " + err.Error())
		return "", echo.NewHTTPError(http.StatusInternalServerError, "could not get access token Spotify access token response body:", err)
	} else {
		return accessToken.(string), nil
	}
}
