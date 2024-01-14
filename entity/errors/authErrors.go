package errors

type WrongPasswordError struct {
	Code    int
	Message string
}

func (e WrongPasswordError) Error() string {
	return e.Message
}
