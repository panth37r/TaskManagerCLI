package user

const ErrAlreadyExists ErrUserAlreadyExists = "user already exists"

type ErrUserAlreadyExists string

func (e ErrUserAlreadyExists) Error() string {
	return string(e)
}
