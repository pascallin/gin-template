package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type VerifyController struct{}

func (a VerifyController) AuthOnly(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	if username == "pascal" && password == "lin" {
		ctx.JSON(http.StatusOK, gin.H{"status": "okay"})
		return
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "failed"})
	}
}
