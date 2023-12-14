package app

import (
	"github.com/gin-gonic/gin"
	"github.com/josephrockqz/wemusic-golang/internal/services"
)

func Run() {
	router := gin.Default()

	router.GET("/spotify-user-authorization-callback", services.SpotifyUserAuthorizationCallback)

	router.Run("localhost:8080")
}
