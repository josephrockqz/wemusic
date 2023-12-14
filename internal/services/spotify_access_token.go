package services

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

// https://developer.spotify.com/documentation/web-api/tutorials/code-flow
func GetAccessToken(code string) (string, error) {
	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", nil)
	if err != nil {
		fmt.Printf("client: could not create Spotify access token request: %s\n", err)
		return "", err
	}

	clientId := os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")
	if clientId == "" || clientSecret == "" {
		fmt.Println("Could not get Spotify client id or client secret. Please set SPOTIFY_CLIENT_ID and SPOTIFY_CLIENT_SECRET environment variables.")
		return "", errors.New("Could not get Spotify client id or client secret from environment")
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
		fmt.Printf("client: could not get response for Spotify access token request: %s\n", err)
		return "", err
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("client: could not read response body for Spotify access token: %s\n", err)
		return "", err
	}

	// TODO: convert response body JSON to map

	// TODO: add error handling for respBody
	// e.g. if "error" key exists in response body map, then return "", error_description value
	// response body: {"access_token":"BQCmET6xTmFN-bryFanwbQJDtSW8TaYwdptSQ8n4wK3HVP9WVMF6QrT-B6mvemjNaP0Y-lfslH44P82nu9fmCz-7w-_wSmsCuQmgjJZDr5ABSBMDfOb2prBVs29aZbLXAk3pjJFSjA6E_OKSlIz1l3kgx0VpM4_5m4cf_oQ-1gVRoxDPf8PbO9AYZ0V8GoK2lA4","token_type":"Bearer","expires_in":3600,"refresh_token":"AQAAuBGB-iejtq3r45-hQOf6i6PlBBiUnwHqQl8sixJl7kGQbj65fQffdWtyhBufA7EBwK3Ak6mmP4MimQzRECDx-khdAQdbP3zWv_T2-sBlIDCBDOuNy9u2e6pZ-pY-0Xc","scope":"user-read-email user-read-private"}
	// response body: {"error":"unsupported_grant_type","error_description":"grant_type parameter is missing"}
	fmt.Println("response body:", string(respBody))

	return "Access Tokennnnnnnn", nil
}
