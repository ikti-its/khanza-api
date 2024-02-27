package exception

type BadRequestError struct {
	Message string
}

type UnauthorizedError struct {
	Message string
}

type ForbiddenError struct {
	Message string
}

type NotFoundError struct {
	Message string
}

type InternalServerError struct {
	Message string
}

func (e BadRequestError) Error() string {
	return e.Message
}

func (e UnauthorizedError) Error() string {
	return e.Message
}

func (e ForbiddenError) Error() string {
	return e.Message
}

func (e NotFoundError) Error() string {
	return e.Message
}

func (e InternalServerError) Error() string {
	return e.Message
}

func PanicIfError(e error) {
	if e != nil {
		switch e.Error() {
		case "duplicated key not allowed":
			panic(BadRequestError{
				Message: "Data already exists",
			})
		case "record not found":
			panic(NotFoundError{
				Message: "Data not found",
			})
		default:
			panic(e)
		}
	}
}
