package constants

const (
	ErrInvalidRequestBody  = "invalid request body"
	ErrInternalServer      = "internal server error"
	ErrInvalidEmail        = "invalid email ID. It should be in the format abcdefg@doamin.com"
	ErrInvalidPhone        = "invalid phone Number. It should be of 10 digits in the format 9876543210."
	ErrInvalidPAN          = "invalid PAN format. It should be in the format ABCDE1234F."
	ErrInvalidPassword     = "invalid password format. It should contain atleast one uppercae , one lowercase and a special character"
	ErrPasswordMismatch    = "password mismatched , password and confirmpassword must match"
	ErrUserExists          = "user already exists"
	ErrCreateUser          = "failed to create user"
	ErrInvalidCredentials  = "invalid email or password"
	ErrInvalidAuthPassword = "invalid password"
	ErrUserNotFound        = "user not registered"
	ErrEmailRequired       = "email is required"
	ErrPasswordRequired    = "password is required"
)

var ErrFieldRequired = map[string]string{
	"name":            "name is required",
	"password":        "password is required",
	"confirmpassword": "confirm Password is required",
	"email":           "email is required",
	"phoneno":         "phone number is required",
	"pan":             "PAN number is required",
}
