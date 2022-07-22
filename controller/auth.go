package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/pascallin/gin-template/service"
)

type AuthController struct{}

type UsernameAndPasswordRequest struct {
	Username string `json:"username" form:"username" xml:"username" binding:"required"`
	Password string `json:"password" form:"password" xml:"password" binding:"required"`
}
type RegisterRequest struct {
	UsernameAndPasswordRequest
	Nickname string `json:"nickname" form:"nickname"`
}
type PatchPasswordRequest struct {
	UsernameAndPasswordRequest
	NewPassword string `json:"newPassword" form:"newPassword" binding:"required"`
}

// @Summary user register
// @Description user register
// @Tags user
// @Accept  json
// @Param user body UsernameAndPasswordRequest true "register"
// @Produce  json
// @Success 200 {object} model.User
// @Router /user/register [post]
func (a AuthController) RegisterRoute(ctx *gin.Context) {
	var request RegisterRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err, id := service.CreteUser(request.Username, request.Password, request.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

// @Summary user login
// @Description user login
// @Tags user
// @Accept  json
// @Param user body UsernameAndPasswordRequest true "login"
// @Produce  json
// @Success 200 {object} model.User
// @Router /user/login [post]
func (a AuthController) LoginRoute(ctx *gin.Context) {
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	fmt.Println(string(body))
	var request UsernameAndPasswordRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err, token := service.Login(request.Username, request.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

// @Summary user patch password
// @Description user patch password
// @Tags user
// @Accept  json
// @Param user body PatchPasswordRequest true "login"
// @Produce  json
// @Success 200 {object} model.User
// @Router /user/password [patch]
func (a AuthController) PatchPasswordRoute(ctx *gin.Context) {
	var request PatchPasswordRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := service.UpdateUserPassword(request.Username, request.Password, request.NewPassword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
