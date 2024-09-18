package response

const (
	ErrorCodeSuccess      = 2001 // success
	ErrorCodeParamInvalid = 2003 // params is invalid
)

// message
var msg = map[int]string{
	ErrorCodeSuccess:      "success",
	ErrorCodeParamInvalid: "param is invalid",
}
