package services

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/josephrockqz/wemusic-golang/internal/transport"
	"github.com/josephrockqz/wemusic-golang/pkg/utils"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func GetAccessToken(context echo.Context, code string) (string, error) {
	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", nil)
	if err != nil {
		zap.L().Error("could not create Spotify access token request: " + err.Error())
		return "", echo.NewHTTPError(http.StatusInternalServerError, "could not create Spotify access token request:", err)
	}

	spotifyClientId, err := utils.GetEnvironmentVariableByName("SPOTIFY_CLIENT_ID")
	if err != nil {
		zap.L().Error("Error retrieving Spotify Client ID from config file")
		return "", echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid Spotify Client ID")
	}
	spotifyClientSecret, err := utils.GetEnvironmentVariableByName("SPOTIFY_CLIENT_SECRET")
	if err != nil {
		zap.L().Error("Error retrieving Spotify Client Secret from config file")
		return "", echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid Spotify Client ID")
	}
	if spotifyClientId == "" || spotifyClientSecret == "" {
		zap.L().Error("Could not get Spotify client id or client secret. Please set SPOTIFY_CLIENT_ID and SPOTIFY_CLIENT_SECRET environment variables.")
		return "", echo.NewHTTPError(http.StatusInternalServerError, "Could not get Spotify client id or client secret from environment")
	}

	query := req.URL.Query()
	query.Add("grant_type", "authorization_code")
	query.Add("code", code)
	query.Add("redirect_uri", "http://localhost:8080/spotify-user-authorization-callback")
	req.URL.RawQuery = query.Encode()

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	clientCredentials := spotifyClientId + ":" + spotifyClientSecret
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

	type AccessTokenResponseSuccessData struct {
		AccessToken  string `json:"access_token"`
		TokenType    string `json:"token_type"`
		ExpiresIn    int    `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
	}

	var accessTokenResponseData AccessTokenResponseSuccessData
	err = json.Unmarshal(respBody, &accessTokenResponseData)

	if err != nil {
		zap.L().Error("could not unmarshal response body for Spotify access token: " + err.Error())
		return "", echo.NewHTTPError(http.StatusInternalServerError, "could not unmarshal response body for Spotify access token:", err)
	}

	accessToken := accessTokenResponseData.AccessToken
	zap.L().Info("access token: " + accessToken)

	accessTokenCookie := &http.Cookie{
		Name:     "spotify_authorize_access_token",
		Value:    accessToken,
		Expires:  time.Now().Add(time.Duration(accessTokenResponseData.ExpiresIn) * time.Second),
		HttpOnly: true, // prevent XSS
	}
	context.SetCookie(accessTokenCookie)

	return accessToken, nil
}
