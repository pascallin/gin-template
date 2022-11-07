package controller

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type FileController struct{}

// @Summary UploadFile
// @Description UploadFile
// @Tags file
// @Accept multipart/form-data
// @Param file formData file true "file"
// @Produce  json
// @Router /upload [post]
func (f FileController) UploadFile(c *gin.Context) {
	// single file
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	// dst.
	dst := fmt.Sprintf("%s.xlsx", time.Now().Format("2006-01-02-15-04-03"))

	c.SaveUploadedFile(file, dst)

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
