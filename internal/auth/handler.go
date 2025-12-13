package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/lukiriskigumilar/resepify-be/pkg/utils"
)

type AuthHandler struct {
	service AuthService
}

func NewAuthHandler(service AuthService) *AuthHandler {
	return &AuthHandler{service}
}

func (h *AuthHandler) RegisterUser(c *gin.Context) {
	var req RegisterRequestDTO

	//validate request body
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.NewApiResponseError(c, "Invalid request", 400, gin.H{
			"reason": err.Error(),
		})
		return
	}

	// call service
	user, err := h.service.RegisterService(req)
	if err != nil {
		utils.NewApiResponseError(c, "Failed to register user", 409, gin.H{
			"reason": err.Error(),
		})
		return
	}

	responseData := RegisterResponseDTO{
		ID:        user.ID.String(),
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	utils.NewApiResponseSuccess(c, "User registered successfully", responseData, 201)
}
