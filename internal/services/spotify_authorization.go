package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/josephrockqz/wemusic-golang/internal/status"
	"github.com/josephrockqz/wemusic-golang/pkg/app"
)

func SpotifyUserAuthorization(context *gin.Context) {
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

	// curl -X GET "https://accounts.spotify.com/authorize?client_id=5fe01282e94241328a84e7c5cc169165&response_type=code&redirect_uri=https%3A%2F%2Fexample.com%2Fcallback&scope=user-read-private%20user-read-email&state=34fFs29kd09"

	appG := app.Gin{C: context}
	context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	context.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	context.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	appG.Response(http.StatusOK, status.SUCCESS, map[string]interface{}{
		"message": "spotifyUserAuthorization",
	})
}
