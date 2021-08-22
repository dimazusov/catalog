package api_error

type Error struct {
	Message string `json:"error"`
}

func New(err error) *Error {
	return &Error{Message: err.Error()}
}
