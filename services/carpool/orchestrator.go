package carpool

import (
	"github.com/imarcelolz/transportation-challenge/services/carpool/enum"
	"github.com/imarcelolz/transportation-challenge/services/carpool/errors"
	"github.com/imarcelolz/transportation-challenge/services/carpool/index"
	"github.com/imarcelolz/transportation-challenge/services/carpool/models"
	"github.com/imarcelolz/transportation-challenge/services/carpool/storage"
	"github.com/imarcelolz/transportation-challenge/services/carpool/tools"
)

type orchestrator struct {
	journeyQueue storage.JourneyIndex
	storage      *storage.Storage
}

func NewOrchestrator() *orchestrator {
	return &orchestrator{
		journeyQueue: index.NewIndex[*models.Journey, int, int](storage.JourneyId, storage.JourneyAvailableSeats),
		storage:      storage.NewStorage(),
	}
}

func (o *orchestrator) Dropoff(id int) error {
	journey := o.storage.FindJourney(id)

	if journey == nil {
		return &errors.ErrorJourneyNotFound{}
	}

	o.storage.DeleteJourney(journey.Id)
	o.processQueue()

	return nil
}

func (o *orchestrator) Locate(id int) (*models.Car, enum.JourneyState) {
	var journey *models.Journey

	if journey = o.storage.FindJourney(id); journey != nil {
		return journey.Car, enum.JourneyInProgress
	}

	if journey = o.journeyQueue.Get(id); journey != nil {
		return journey.Car, enum.JourneyQueued
	}

	return nil, enum.JourneyUndefined
}

func (o *orchestrator) RegisterJourney(apiJourney models.Journey) error {
	journey := models.NewJourney(apiJourney.Id, apiJourney.People)

	if _, state := o.Locate(journey.Id); state != enum.JourneyUndefined {
		return &errors.ErrorJourneyAlreadyExists{}
	}

	o.journeyQueue.Add(journey)
	o.processQueue()

	return nil
}

func (o *orchestrator) RegisterCars(apiCars []models.Car) error {
	if !tools.IsUniqueBy(apiCars, func(car models.Car) int { return car.Id }) {
		return &errors.ErrorDuplicatedCarId{}
	}

	for _, car := range apiCars {
		o.storage.AddCar(models.NewCar(car.Id, car.Seats))
	}

	return nil
}

func (o *orchestrator) processQueue() {
	var journey *models.Journey
	var car *models.Car

	for _, key := range o.journeyQueue.Keys() {
		if journey = o.journeyQueue.Get(key); journey == nil {
			continue
		}

		if car = o.storage.FindCarByRequiredSeats(journey.People); car == nil {
			continue
		}

		o.journeyQueue.Remove(journey)
		o.storage.AddJourney(journey, car)
	}
}
