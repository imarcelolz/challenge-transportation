package storage

import (
	"github.com/imarcelolz/transportation-challenge/services/carpool/index"
	"github.com/imarcelolz/transportation-challenge/services/carpool/models"
)

type Storage struct {
	cars     CarIndex
	journeys JourneyIndex
}

func NewStorage() *Storage {
	return &Storage{
		cars:     index.NewIndex[*models.Car, int, int](CarId, CarAvailableSeats),
		journeys: index.NewIndex[*models.Journey, int, int](JourneyId, JourneyAvailableSeats),
	}
}

func (s *Storage) AddCar(car *models.Car) {
	s.cars.Add(car)
}

func (s *Storage) AddJourney(journey *models.Journey, car *models.Car) {
	s.cars.Remove(car)
	journeys := s.removeJourneys(car.JourneyIds)

	car.AddJourney(journey)

	s.cars.Add(car)
	s.journeys.Add(journey)
	s.addJourneys(journeys)
}

func (s *Storage) DeleteJourney(id int) {
	var journey *models.Journey

	if journey = s.journeys.Get(id); journey == nil {
		return
	}

	car := journey.Car

	s.cars.Remove(car)
	s.journeys.Remove(journey)

	journeys := s.removeJourneys(car.JourneyIds)
	car.RemoveJourney(journey)

	s.cars.Add(car)
	s.addJourneys(journeys)
}

func (s *Storage) FindCar(id int) *models.Car {
	return s.cars.Get(id)
}

func (s *Storage) FindJourney(id int) *models.Journey {
	return s.journeys.Get(id)
}

func (s *Storage) FindCarByRequiredSeats(requiredSeats int) *models.Car {
	for _, key := range s.cars.SecondaryKeys() {
		if key < requiredSeats {
			continue
		}

		if cars := s.cars.SecondaryIndex(key); len(cars) > 0 {
			return cars[0]
		}
	}

	return nil
}

func (s *Storage) addJourneys(journeys []*models.Journey) {
	for _, journey := range journeys {
		s.journeys.Add(journey)
	}
}

func (s *Storage) removeJourneys(journeyIds []int) []*models.Journey {
	journeys := make([]*models.Journey, 0, len(journeyIds))

	for _, journeyId := range journeyIds {
		journey := s.journeys.Get(journeyId)

		if journey != nil {
			journeys = append(journeys, journey)
			s.journeys.Remove(journey)
		}
	}

	return journeys
}
