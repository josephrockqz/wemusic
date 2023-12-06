package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/josephrockqz/wemusic-golang/internal/services"
)

func Run() {
	fmt.Println("app init")

	router := gin.Default()
	router.POST("/access-token", services.GetAccessToken)
	router.GET("/spotify-user-authorization", services.SpotifyUserAuthorization)
	router.GET("/callback", services.SpotifyLoginCallback)

	router.Run("localhost:8080")
}
