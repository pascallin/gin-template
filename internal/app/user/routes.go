package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/user")
	router.POST("/register", registerRoute)
	router.POST("/login", loginRoute)
	router.PATCH("/password", patchPasswordRoute)
}

type UsernameAndPasswordRequest struct {
	Username string	`json:"username" form:"username" xml:"username" binding:"required"`
	Password string	`json:"password" form:"password" xml:"password" binding:"required"`
}
type RegisterRequest struct {
	UsernameAndPasswordRequest
	Nickname string	`json:"nickname" form:"nickname"`
}
type PatchPasswordRequest struct {
	UsernameAndPasswordRequest
	NewPassword string	`json:"newPassword" form:"newPassword" binding:"required"`
}
func registerRoute(ctx *gin.Context) {
	var request RegisterRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err, id := register(request.Username, request.Password, request.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"id":id})
}

func loginRoute(ctx *gin.Context) {
	var request UsernameAndPasswordRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err, token := login(request.Username, request.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token":token})
}

func patchPasswordRoute(ctx *gin.Context) {
	var request PatchPasswordRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := updatePassword(request.Username, request.Password, request.NewPassword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status":"success"})
}