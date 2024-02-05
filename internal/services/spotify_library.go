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
) ([]byte, int, error) {
	req, err := http.NewRequest(method, "https://api.spotify.com/"+endpoint, body)
	if err != nil {
		zap.L().Error("could not create Spotify library request: " + err.Error())
		return nil, 0, err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := transport.HttpClient()

	resp, err := client.Do(req)
	if err != nil {
		zap.L().Error("could not get response for Spotify library request: " + err.Error())
		return nil, 0, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		zap.L().Error("could not read response body for Spotify library: " + err.Error())
		return nil, 0, err
	}

	return respBody, resp.StatusCode, nil
}

func GetProfile(context echo.Context, accessToken string) error {
	respBody, statusCode, err := fetchWebApi("v1/me", "GET", nil, accessToken)
	if err != nil {
		zap.L().Error("could not fetch Spotify web API: " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "could not fetch Spotify web API:", err)
	}
	bodyString := string(respBody)
	zap.L().Info(bodyString)
	zap.L().Sugar().Infof("Status Code: %d", statusCode)
	if statusCode != 200 {
		zap.L().Error("Spotify top tracks request returned with error code")
		return echo.NewHTTPError(http.StatusInternalServerError, "Spotify top tracks request returned with error code")
	}

	var responseData models.TopTracksResponseData
	// err = json.NewDecoder(resp.Body).Decode(&responseData)
	err = json.Unmarshal(respBody, &responseData)
	if err != nil {
		zap.L().Error("could not unmarshal response body for Spotify top songs: " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "could not unmarshal response body for Spotify library:", err)
	}

	zap.L().Info(responseData.Href)
	// zap.L().Info(responseData.Items[0].Name + ", " + responseData.Items[0].Type)

	return nil
}
