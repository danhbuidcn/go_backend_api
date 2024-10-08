package response

const (
	ErrorCodeSuccess      = 2001 // success
	ErrorCodeParamInvalid = 2003 // params is invalid

	ErrorInvalidToken = 3001 // token is invalid
	ErrInvalidOTP     = 3002
	ErrSendEmailOtp   = 3003
	// Register Code
	ErrCodeUserHasExists = 5001 // user has already registered
)

// message
var msg = map[int]string{
	ErrorCodeSuccess:      "success",
	ErrorCodeParamInvalid: "param is invalid",
	ErrorInvalidToken:     "token is invalid",
	ErrInvalidOTP:         "Otp error",
	ErrSendEmailOtp:       "Failed to send email OTP",

	ErrCodeUserHasExists: "user has already registerd",
}
