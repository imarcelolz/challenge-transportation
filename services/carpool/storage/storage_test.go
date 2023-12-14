package storage

import (
	"testing"

	"github.com/imarcelolz/transportation-challenge/services/carpool/models"
	"github.com/stretchr/testify/assert"
)

func TestAddCar(t *testing.T) {
	t.Run("adds a car to the storage", func(t *testing.T) {
		storage := NewStorage()
		storage.AddCar(models.NewCar(1, 5))

		assert.Equal(t, storage.cars.Get(1).Id, 1)
	})
}

func TestAddJourney(t *testing.T) {
	t.Run("adds a journey to the storage", func(t *testing.T) {
		storage := NewStorage()
		car, journey := models.NewCar(1, 5), models.NewJourney(1, 4)

		storage.AddCar(car)
		storage.AddJourney(journey, car)

		assert.Equal(t, storage.journeys.Get(1), journey)
	})

	t.Run("reindex the car", func(t *testing.T) {
		storage := NewStorage()
		car, journey := models.NewCar(1, 5), models.NewJourney(1, 4)

		storage.AddCar(car)
		storage.AddJourney(journey, car)

		assert.Equal(t, storage.cars.Get(1).AvailableSeats, 1)
	})

	t.Run("reindex the previous journeys car", func(t *testing.T) {
		storage := NewStorage()
		car, journey := models.NewCar(1, 10), models.NewJourney(1, 5)
		secondJourney := models.NewJourney(2, 3)
		thirdJourney := models.NewJourney(3, 1)
		fourthJourney := models.NewJourney(4, 1)
		journeys := []*models.Journey{journey, secondJourney, thirdJourney, fourthJourney}

		storage.AddCar(car)
		for _, j := range journeys {
			storage.AddJourney(j, car)
		}

		secondaryIndex := storage.journeys.SecondaryIndex(0)
		for _, j := range journeys {
			assert.Equal(t, storage.journeys.Get(j.Id), j)
			assert.Contains(t, secondaryIndex, j)
			assert.Contains(t, car.JourneyIds, j.Id)
		}
	})
}

func TestDeleteJourney(t *testing.T) {
	t.Run("removes a journey from the storage", func(t *testing.T) {
		storage := NewStorage()
		car, journey := models.NewCar(1, 10), models.NewJourney(1, 5)

		storage.AddCar(car)
		storage.AddJourney(journey, car)
		storage.DeleteJourney(1)

		assert.Nil(t, storage.journeys.Get(1))
		assert.Empty(t, storage.journeys.SecondaryIndex(2))
		assert.Empty(t, car.JourneyIds)
	})

	t.Run("reindex the car", func(t *testing.T) {
		car, journey := models.NewCar(1, 10), models.NewJourney(1, 5)
		secondJourney := models.NewJourney(2, 3)
		thirdJourney := models.NewJourney(3, 1)
		fourthJourney := models.NewJourney(4, 1)
		journeys := []*models.Journey{journey, secondJourney, thirdJourney, fourthJourney}

		storage := NewStorage()
		storage.AddCar(car)

		for _, j := range journeys {
			storage.AddJourney(j, car)
		}

		storage.DeleteJourney(1)

		assert.Equal(t, storage.cars.Get(1).AvailableSeats, 5)
		assert.Contains(t, car.JourneyIds, 2)
		assert.Contains(t, car.JourneyIds, 3)
		assert.Contains(t, car.JourneyIds, 4)
	})

	t.Run("reindex the previous journeys car", func(t *testing.T) {
		car, journey := models.NewCar(1, 10), models.NewJourney(1, 5)
		secondJourney := models.NewJourney(2, 3)
		thirdJourney := models.NewJourney(3, 1)
		fourthJourney := models.NewJourney(4, 1)
		journeys := []*models.Journey{journey, secondJourney, thirdJourney, fourthJourney}

		storage := NewStorage()
		storage.AddCar(car)

		for _, j := range journeys {
			storage.AddJourney(j, car)
		}

		assert.Equal(t, car.AvailableSeats, 0)
		assert.Equal(t, storage.FindCar(1).AvailableSeats, 0)

		secondaryIndex := storage.journeys.SecondaryIndex(0)
		for _, j := range journeys {
			assert.Contains(t, secondaryIndex, j)
		}

		storage.DeleteJourney(1)
		storage.DeleteJourney(2)

		assert.Equal(t, car.AvailableSeats, 8)
		assert.Equal(t, storage.FindCar(1).AvailableSeats, 8)

		secondaryIndex = storage.journeys.SecondaryIndex(8)

		for _, j := range []*models.Journey{journey, secondJourney} {
			assert.NotContains(t, secondaryIndex, j)
			assert.NotContains(t, car.JourneyIds, j.Id)
		}

		for _, j := range []*models.Journey{thirdJourney, fourthJourney} {
			assert.Contains(t, secondaryIndex, j)
			assert.Contains(t, car.JourneyIds, j.Id)
		}
	})
}

func TestFindCar(t *testing.T) {
	t.Run("returns the car", func(t *testing.T) {
		storage := NewStorage()
		car := models.NewCar(1, 5)

		storage.AddCar(car)

		assert.Equal(t, storage.FindCar(1), car)
	})

	t.Run("returns nil when the car does not exist", func(t *testing.T) {
		storage := NewStorage()
		assert.Nil(t, storage.FindCar(1))
	})
}

func TestFindJourney(t *testing.T) {
	t.Run("returns the journey", func(t *testing.T) {
		storage := NewStorage()
		car, journey := models.NewCar(1, 5), models.NewJourney(1, 4)

		storage.AddCar(car)
		storage.AddJourney(journey, car)

		assert.Equal(t, storage.FindJourney(1), journey)
	})
	t.Run("returns nil when the journey does not exist", func(t *testing.T) {
		storage := NewStorage()
		assert.Nil(t, storage.FindJourney(1))
	})
}

func TestFindCarByRequiredSeats(t *testing.T) {
	t.Run("returns the car", func(t *testing.T) {
		storage := NewStorage()
		car := models.NewCar(1, 5)

		storage.AddCar(car)

		assert.Equal(t, storage.FindCarByRequiredSeats(5), car)
	})

	t.Run("returns nil when the car does not exist", func(t *testing.T) {
		storage := NewStorage()
		assert.Nil(t, storage.FindCarByRequiredSeats(5))
	})

	t.Run("returns the car with the minimum available seats", func(t *testing.T) {
		storage := NewStorage()
		car := models.NewCar(1, 5)
		secondCar := models.NewCar(2, 10)

		storage.AddCar(car)
		storage.AddCar(secondCar)

		assert.Equal(t, storage.FindCarByRequiredSeats(5), car)
	})
}
