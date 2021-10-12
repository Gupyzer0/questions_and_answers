package utils

/*import(
	"errors"
)*/

type CustomError struct {
	message string
}

func (err CustomError) Error() string {
	return err.message
}

var ErrNotFound = CustomError{message: "Resource Not Found"}
var ErrBadData = CustomError{message: "Bad Request"}
var ServerError = CustomError{message: "Server Error"}

/*var (
	//ErrNotFound = errors.New("Not found")
	ErrBadData = errors.New("Bad data")
)*/
