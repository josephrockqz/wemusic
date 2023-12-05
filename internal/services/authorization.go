package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAccessToken(c *gin.Context) {
	// curl -X POST -H "Content-Type: application/x-www-form-urlencoded" -d 'grant_type=client_credentials&client_id=YOUR_CLIENT_ID&client_secret=YOUR_CLIENT_SECRET' "https://accounts.spotify.com/api/token"
	// https://developer.spotify.com/documentation/general/guides/authorization-guide/#client-credentials-flow
	// https://developer.spotify.com/console/get-search-item/?q=tania%20bowra&type=artist&market=&limit=&offset=
	// https://developer.spotify.com/documentation/web-api/reference/#category-search
	// https://developer.spotify.com/documentation/web-api/reference/#endpoint-search
	c.IndentedJSON(http.StatusOK, gin.H{"message": "getAccessToken"})
}
