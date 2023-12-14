package storage

import (
	"github.com/imarcelolz/transportation-challenge/services/carpool/index"
	"github.com/imarcelolz/transportation-challenge/services/carpool/models"
)

type CarIndex = *index.Index[*models.Car, int, int]
type JourneyIndex = *index.Index[*models.Journey, int, int]
