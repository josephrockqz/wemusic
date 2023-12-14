package services

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/josephrockqz/wemusic-golang/internal/status"
	"github.com/josephrockqz/wemusic-golang/pkg/app"
)

// https://developer.spotify.com/documentation/web-api/tutorials/code-flow
func GetAccessToken(context *gin.Context) {
	appG := app.Gin{C: context}
	// context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	// context.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	// context.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", nil)
	if err != nil {
		fmt.Printf("client: could not create Spotify access token request: %s\n", err)
		appG.Response(http.StatusServiceUnavailable, status.ERROR, map[string]interface{}{
			"message": "spotifyAccessToken",
		})
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	clientId := os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")
	if clientId == "" || clientSecret == "" {
		fmt.Println("Could not get Spotify client id or client secret. Please set SPOTIFY_CLIENT_ID and SPOTIFY_CLIENT_SECRET environment variables.")
		appG.Response(http.StatusServiceUnavailable, status.ERROR, map[string]interface{}{
			"message": "could not get Spotify client id or client secret",
		})
		return
	}

	query := req.URL.Query()
	query.Add("grant_type", "client_credentials")
	query.Add("client_id", clientId)
	query.Add("client_secret", clientSecret)
	req.URL.RawQuery = query.Encode()

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: could not get response for Spotify access token: %s\n", err)
		appG.Response(http.StatusServiceUnavailable, status.ERROR, map[string]interface{}{
			"message": "could not get response for Spotify access token",
		})
		return
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("client: could not read response body for Spotify access token: %s\n", err)
		appG.Response(http.StatusServiceUnavailable, status.ERROR, map[string]interface{}{
			"message": "could not read response body for Spotify access token",
		})
		return
	}
	fmt.Println(string(respBody))

	appG.Response(http.StatusOK, status.SUCCESS, map[string]interface{}{
		"message": "spotifyAccessToken",
		"data":    string(respBody),
	})
}
