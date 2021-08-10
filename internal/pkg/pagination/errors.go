package pagination

import "github.com/pkg/errors"

var ErrExceededMaxCountPerPage = errors.New("Exceeded max count per page")
var ErrWrongCountPerPage = errors.New("Wrong count per page")
var ErrWrongPage = errors.New("Wrong page")


