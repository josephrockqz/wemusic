package services

import (
	// "bytes"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SpotifyUserAuthorization(c *gin.Context) {
	// client := &http.Client{
	// 	CheckRedirect: redirectPolicyFunc,
	// }

	// jsonBody := []byte(`{"client_id":"5fe01282e94241328a84e7c5cc169165","response_type":"code","redirect_uri":"https://example.com/callback","scope":"user-read-private user-read-email","state":"34fFs29kd09"}`)
	// bodyReader := bytes.NewReader(jsonBody)

	// req, err := http.NewRequest("GET", "https://accounts.spotify.com/authorize", bodyReader)
	// if err != nil {
	// 	fmt.Printf("client: could not create request for Spotify user authorization: %s\n", err)
	// 	return
	// }

	fmt.Println("SpotifyUserAuthorization")
	c.IndentedJSON(http.StatusOK, gin.H{"message": "spotifyUserAuthorization"})

	// curl -X GET "https://accounts.spotify.com/authorize?client_id=5fe01282e94241328a84e7c5cc169165&response_type=code&redirect_uri=https%3A%2F%2Fexample.com%2Fcallback&scope=user-read-private%20user-read-email&state=34fFs29kd09"
}

func SpotifyLoginCallback(c *gin.Context) {
	fmt.Println("SpotifyLoginCallback")
	c.IndentedJSON(http.StatusOK, gin.H{"message": "spotifyLoginCallback"})
}

func GetAccessToken(c *gin.Context) {
	// curl -X POST -H "Content-Type: application/x-www-form-urlencoded" -d 'grant_type=client_credentials&client_id=YOUR_CLIENT_ID&client_secret=YOUR_CLIENT_SECRET' "https://accounts.spotify.com/api/token"
	// https://developer.spotify.com/documentation/general/guides/authorization-guide/#client-credentials-flow
	// https://developer.spotify.com/console/get-search-item/?q=tania%20bowra&type=artist&market=&limit=&offset=
	// https://developer.spotify.com/documentation/web-api/reference/#category-search
	// https://developer.spotify.com/documentation/web-api/reference/#endpoint-search
	c.IndentedJSON(http.StatusOK, gin.H{"message": "getAccessToken"})
}
