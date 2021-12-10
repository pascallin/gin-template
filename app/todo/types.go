package todo

import (
	"github.com/pascallin/gin-template/pkg"
)

type Todo struct {
	pkg.GormModel
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CreateTodoInput struct {
	Title       string `form:"title" xml:"title" json:"title" binding:"required"`
	Description string `json:"description"`
}

type UpdateTodoInput struct {
	ID          uint64 `uri:"id" binding:"required" json:"id"`
	Title       string `form:"title" xml:"title" json:"title"`
	Description string `form:"title" xml:"title" json:"description"`
}

type Pagination struct {
	PageSize uint64 `form:"pageSize" binding:"required,max=20"`
	Page     uint64 `form:"page" binding:"required,max=100"`
}
