package app

import (
	"github.com/gin-gonic/gin"
	"github.com/josephrockqz/wemusic-golang/internal/services"
	"github.com/josephrockqz/wemusic-golang/internal/transport/middleware"
)

func Run() {
	router := gin.Default()

	// router.SetTrustedProxies([]string{"127.0.0.1"})

	router.Use(middleware.CorsMiddleware())

	router.GET("/spotify-login", services.SpotifyLogin)
	router.GET("/spotify-user-authorization-callback", services.SpotifyUserAuthorizationCallback)

	router.Run("localhost:8080")
}
