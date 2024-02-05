package response

import "errors"

var (
	ErrBookNotFound         = errors.New("book not found")
	ErrBookNameAlreadyTaken = errors.New("book name already taken")
	ErrNameEmpty            = errors.New("name is empty")
)
