package tasks

const ErrTaskNotFound ErrNotFound = "cant find task by id"
const ErrUnknownStatusCode ErrUnknownCode = "unknown status code"

type ErrUnknownCode string

func (e ErrUnknownCode) Error() string {
	return string(e)
}

type ErrNotFound string

func (e ErrNotFound) Error() string {
	return string(e)
}
