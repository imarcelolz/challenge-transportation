package errors

import "fmt"

type ErrorJourneyAlreadyExists struct {
	id int
}

func (e ErrorJourneyAlreadyExists) Error() string {
	return fmt.Sprintf("Journey ID %d already exists", e.id)
}
