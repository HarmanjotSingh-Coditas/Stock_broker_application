package models

type BFFUserRequest struct {
	UserName        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirmpassword" binding:"required" validate:"eqfield=ConfirmPassword"`
	EmailId         string `json:"email" binding:"required"`
	PhoneNumber     uint64 `json:"phonenumber" binding:"required"`
	PANNumber       string `json:"pan" binding:"required"`
}

type BFFUserResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
