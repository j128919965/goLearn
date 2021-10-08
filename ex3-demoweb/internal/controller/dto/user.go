package dto

type CreateUserRequest struct {
	Id   int    `json:"id" form:"id" binding:"required"`
	Name string `json:"name" form:"name" binding:"required"`
	Age  *int    `json:"age" form:"age"`
}

type GetUserResponse struct {
	Id   int    `json:"id" `
	Name string `json:"name"`
	Age  *int    `json:"age" `
}