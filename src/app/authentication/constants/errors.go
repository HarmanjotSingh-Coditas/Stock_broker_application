package constants

const (
	ErrInvalidRequestBody = "Invalid request body"
	ErrInternalServer     = "Internal server error"
	ErrInvalidEmail       = "Invalid Email ID. It should be in the format abcdefg@doamin.com"
	ErrInvalidPhone       = "Invalid Phone Number. It should be of 10 digits in the format 9876543210."
	ErrInvalidPAN         = "Invalid PAN format. It should be in the format ABCDE1234F."
	ErrInvalidPassword    = "Invalid Password format. It should contain atleast one uppercae , one lowercase"
	ErrPasswordMismatch   = "Password Mismatched , password and confirmpassword must match"
	ErrUserExists         = "user already exists"
	ErrCreateUser         = "failed to create user"
)
