package errhandler

type CommonError struct {
	Message string            `json:"message"`
	Code    string            `json:"code"`
	Errors  []ValidationError `json:"errors"`
}

type ValidationError struct {
	Path    string `json:"path"`
	Message string `json:"message"`
}

type RatingNotFoundError CommonError

func (e *RatingNotFoundError) Error() string {
	return e.Message
}

func NewRatingNotFoundError(message string) error {
	return &RatingNotFoundError{
		Message: message,
		Code:    "rating.not.found",
	}
}

type BadRequestError CommonError

func (e *BadRequestError) Error() string {
	return e.Message
}

func NewBadRequestError(message string, errors []ValidationError) error {
	return &BadRequestError{
		Message: message,
		Code:    "invalid.request",
		Errors:  errors,
	}
}

type PermissionDeniedError CommonError

func (e *PermissionDeniedError) Error() string {
	return e.Message
}

func NewPermissionDeniedError(message string) error {
	return &PermissionDeniedError{
		Message: message,
		Code:    "permission.denied",
	}
}
