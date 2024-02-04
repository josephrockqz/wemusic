package services

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/josephrockqz/wemusic-golang/internal/transport"
	"github.com/josephrockqz/wemusic-golang/models"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func fetchWebApi(
	endpoint string,
	method string,
	body io.Reader,
	accessToken string,
) ([]byte, error) {
	req, err := http.NewRequest(method, "https://api.spotify.com/"+endpoint, body)
	if err != nil {
		zap.L().Error("could not create Spotify library request: " + err.Error())
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := transport.HttpClient()

	resp, err := client.Do(req)
	if err != nil {
		zap.L().Error("could not get response for Spotify library request: " + err.Error())
		return nil, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		zap.L().Error("could not read response body for Spotify library: " + err.Error())
		return nil, err
	}

	return respBody, nil
}

func GetLikedSongs(context echo.Context, accessToken string) error {
	respBody, err := fetchWebApi("v1/me", "GET", nil, accessToken)
	if err != nil {
		zap.L().Error("could not fetch Spotify web API: " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "could not fetch Spotify web API:", err)
	}

	zap.L().Info("successfully made library request")

	var responseData models.AccessTokenResponseSuccessData
	err = json.Unmarshal(respBody, &responseData)
	// if err != nil {
	// 	zap.L().Error("could not unmarshal response body for Spotify library: " + err.Error())
	// 	return echo.NewHTTPError(http.StatusInternalServerError, "could not unmarshal response body for Spotify library:", err)
	// }

	return nil
}
