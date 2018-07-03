package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
	ErrBind             = &Errno{Code: 20102, Message: "The user was not found."}

	// user errors
	ErrUserNotFount = &Errno{Code: 20102, Message: "The user was not found."}
)
