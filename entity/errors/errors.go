package errors

type ErrorType struct {
	Code    int
	Message string
}
type WrongPasswordError struct {
	ErrorType
}
type UserNotFoundError struct {
	ErrorType
}
type UserIsNotActiveError struct {
	ErrorType
}

func (e ErrorType) Error() string {
	return e.Message
}
