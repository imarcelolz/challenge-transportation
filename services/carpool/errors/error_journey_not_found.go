package errors

import "fmt"

type ErrorJourneyNotFound struct {
	journeyId int
}

func (e *ErrorJourneyNotFound) Error() string {
	return fmt.Sprintf("Journey with id %d does not exist in the trip", e.journeyId)
}
