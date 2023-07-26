package api

type BadRequestError struct {
	msg string
}

func (e BadRequestError) Error() string {
	return e.msg
}

type NotFoundError struct {
	msg string
}

func (e NotFoundError) Error() string {
	return e.msg
}

type AuthError struct {
	msg string
}

func (e AuthError) Error() string {
	return e.msg
}