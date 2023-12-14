package errors

import "fmt"

type ErrorDuplicatedCarId struct {
	id int
}

func (e ErrorDuplicatedCarId) Error() string {
	return fmt.Sprintf("Car ID %d is duplicated", e.id)
}
