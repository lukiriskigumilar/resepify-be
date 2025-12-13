package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lukiriskigumilar/resepify-be/internal/auth"
)

func GlobalRoutes(api *gin.RouterGroup, authModule *auth.AuthModule) {
	auth.AuthRoutes(api, authModule)
}
