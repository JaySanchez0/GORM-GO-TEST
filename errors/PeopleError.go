package errors

type PeopleError struct {
	error
	msg string
}

func (e *PeopleError) Error() string {
	return e.msg
}

func New(msg string) error {
	return &PeopleError{msg: msg}
}
