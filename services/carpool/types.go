package carpool

import (
	"github.com/imarcelolz/transportation-challenge/services/carpool/enum"
	"github.com/imarcelolz/transportation-challenge/services/carpool/models"
)

type Orchestrator interface {
	Dropoff(id int) error
	Locate(id int) (*models.Car, enum.JourneyState)
	RegisterJourney(journey models.Journey) error
	RegisterCars(cars []models.Car) error
}
