package storage

import "github.com/imarcelolz/transportation-challenge/services/carpool/models"

func CarId(car *models.Car) int {
	return car.Id
}

func CarAvailableSeats(car *models.Car) int {
	return car.AvailableSeats
}

func JourneyId(journey *models.Journey) int {
	return journey.Id
}

func JourneyAvailableSeats(journey *models.Journey) int {
	if journey.Car != nil {
		return journey.Car.AvailableSeats
	}

	return 0
}
