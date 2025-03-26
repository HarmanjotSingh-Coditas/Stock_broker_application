package models

type BFFUserSignInRequest struct {
	EmailId  string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type BFFUserSignInResponse struct {
	Message string `json:"message"`
}
