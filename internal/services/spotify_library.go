package services

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetLibrary(context echo.Context, accessToken string) error {
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/me", nil)
	if err != nil {
		context.Logger().Error("could not create Spotify library request:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "could not create Spotify library request:", err)
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		context.Logger().Error("could not get response for Spotify library request:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "could not get response for Spotify library request:", err)
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		context.Logger().Error("could not read response body for Spotify library:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "could not read response body for Spotify library:", err)
	}

	var accessTokenResponseData map[string]interface{}
	err = json.Unmarshal(respBody, &accessTokenResponseData)
	if err != nil {
		context.Logger().Error("could not unmarshal response body for Spotify library:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "could not unmarshal response body for Spotify library:", err)
	}

	return nil
}
