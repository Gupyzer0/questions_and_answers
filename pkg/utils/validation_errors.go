package utils

type ValidationErrors interface {
	AddError(error)
	GetErrors() []string
}

type CustomValidationErrors struct {
	Errors []error
}

func (err CustomValidationErrors) Error() string {

	var err_msg string

	for _, e := range err.Errors {
		err_msg = err_msg + e.Error() + "\n "
	}

	return err_msg
}

func (err *CustomValidationErrors) AddError(e error) {
	err.Errors = append(err.Errors, e)
}

func (err *CustomValidationErrors) GetErrors() []string {
	string_errors := make([]string, 0)

	for _, e := range err.Errors {
		string_errors = append(string_errors, e.Error())
	}

	return string_errors
}
