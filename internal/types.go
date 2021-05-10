package internal

type Pagination struct {
	PageSize uint64 `form:"pageSize" binding:"required,max=20"`
	Page uint64 `form:"page" binding:"required,max=100"`
}