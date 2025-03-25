package validators

import "regexp"

var (
	PANRegex   = regexp.MustCompile(`^[A-Z]{5}[0-9]{4}[A-Z]{1}$`)
	EmailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
)

func ValidatePAN(pan string) bool {
	return PANRegex.MatchString(pan)
}

func ValidatePhone(phone uint64) bool {
	return phone >= 1000000000 && phone <= 9999999999
}

func ValidateEmail(emailid string) bool {
	return EmailRegex.MatchString(emailid)
}
