package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetLibrary(accessToken string) error {
	fmt.Println("Get Library function:", accessToken)

	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/me", nil)
	if err != nil {
		fmt.Printf("could not create Spotify library request: %s\n", err)
		return err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("could not get response for Spotify library request: %s\n", err)
		return err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("could not read response body for Spotify library: %s\n", err)
		return err
	}

	var accessTokenResponseData map[string]interface{}
	err = json.Unmarshal(respBody, &accessTokenResponseData)
	if err != nil {
		fmt.Printf("could not unmarshal response body for Spotify library: %s\n", err)
		return err
	}

	fmt.Println(accessTokenResponseData)

	return nil
}
