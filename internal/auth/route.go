package auth

import "github.com/gin-gonic/gin"

func AuthRoutes(ar *gin.RouterGroup, m *AuthModule) {
	auth := ar.Group("/auth")

	auth.POST("/register", m.Handler.RegisterUser)
}
