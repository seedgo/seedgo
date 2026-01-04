package seedgo

type BusErr struct {
	Code    int
	Message string
	Err     error
}

func NewBusErr(code int, err error, message string) error {
	return BusErr{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

func (busErr BusErr) Error() string {
	if busErr.Err == nil {
		return busErr.Message
	}
	return busErr.Err.Error()
}

func (busErr BusErr) Unwrap() error {
	return busErr.Err
}

var (
	SystemErr          = BusErr{Code: 400, Message: "system error"}
	UnAuthenticateErr  = BusErr{Code: 401, Message: "unauthentication "}
	UnAuthorizationErr = BusErr{Code: 403, Message: "unauthorization"}
)
