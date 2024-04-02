package exception

import "log"

type (
	BadRequestError     struct{ Message string }
	UnauthorizedError   struct{ Message string }
	ForbiddenError      struct{ Message string }
	NotFoundError       struct{ Message string }
	InternalServerError struct{ Message string }
	ValidationError     struct {
		FailedField string
		Tag         string
		Value       any
	}
)

func (e *BadRequestError) Error() string {
	return e.Message
}
func (e *UnauthorizedError) Error() string {
	return e.Message
}
func (e *ForbiddenError) Error() string {
	return e.Message
}
func (e *NotFoundError) Error() string {
	return e.Message
}
func (e *InternalServerError) Error() string {
	return e.Message
}

func PanicIfError(err error, ctx ...string) {
	if err != nil {
		if err.Error() == "record not found" {
			panic(&NotFoundError{
				Message: "Data not found",
			})
		} else {
			log.Printf("Error: %v %v", ctx, err)
			panic(&InternalServerError{
				Message: "Internal Server Error",
			})
		}
	}
}
