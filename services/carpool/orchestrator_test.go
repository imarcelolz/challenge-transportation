package carpool

import (
	"testing"

	"github.com/imarcelolz/transportation-challenge/services/carpool/enum"
	"github.com/imarcelolz/transportation-challenge/services/carpool/errors"
	"github.com/imarcelolz/transportation-challenge/services/carpool/models"
	"github.com/stretchr/testify/assert"
)

const JOURNEY_ID = 1
const JOURNEY_PEOPLE = 1
const CAR_ID = 1
const CAR_SEATS = 2
const INVALID_JOURNEY_ID = 5

func TestOrchestratorDropoff(t *testing.T) {
	t.Run("when trip is found returns no error", func(t *testing.T) {
		target := createOrchestrator()
		err := target.Dropoff(JOURNEY_ID)
		_, state := target.Locate(JOURNEY_ID)

		assert.Nil(t, err)
		assert.Equal(t, enum.JourneyUndefined, state)
	})

	t.Run("when trip is found reprocess the queue", func(t *testing.T) {
		target := createOrchestrator()
		target.RegisterJourney(*models.NewJourney(2, 1))

		target.Dropoff(JOURNEY_ID)
		_, state := target.Locate(2)

		assert.Equal(t, state, enum.JourneyInProgress)
	})

	t.Run("when trip is not found", func(t *testing.T) {
		target := createOrchestrator()

		err := target.Dropoff(INVALID_JOURNEY_ID)

		assert.IsType(t, &errors.ErrorJourneyNotFound{}, err)
	})

	t.Run("when trip is finished, releases the car", func(t *testing.T) {
		target := createOrchestrator()

		err := target.Dropoff(JOURNEY_ID)

		assert.Nil(t, err)
		assert.Empty(t, len(target.storage.FindCar(CAR_ID).JourneyIds))
	})

	t.Run("when the journey is enqueued, remove it from the queue", func(t *testing.T) {
		target := createOrchestrator()

		target.RegisterJourney(*models.NewJourney(3, CAR_SEATS))
		_, initialState := target.Locate(3)

		err := target.Dropoff(JOURNEY_ID)

		_, state := target.Locate(3)

		assert.Nil(t, err)
		assert.Equal(t, enum.JourneyQueued, initialState)
		assert.Equal(t, enum.JourneyInProgress, state)
	})
}

func TestOrchestratorLocate(t *testing.T) {
	t.Run("when journey is found returns the car", func(t *testing.T) {
		target := createOrchestrator()
		_, state := target.Locate(JOURNEY_ID)

		assert.Equal(t, enum.JourneyInProgress, state)
	})

	t.Run("when journey is queued returns queued state", func(t *testing.T) {
		target := createOrchestrator()
		target.RegisterJourney(*models.NewJourney(3, 10))

		_, state := target.Locate(3)

		assert.Equal(t, enum.JourneyQueued, state)
	})

	t.Run("when journey is not found returns undefined state", func(t *testing.T) {
		target := createOrchestrator()

		_, state := target.Locate(INVALID_JOURNEY_ID)

		assert.Equal(t, enum.JourneyUndefined, state)
	})
}

func TestOrchestratorRegisterJourney(t *testing.T) {
	t.Run("when journey is valid returns no error", func(t *testing.T) {
		target := createOrchestrator()
		err := target.RegisterJourney(*models.NewJourney(2, 1))

		_, state := target.Locate(2)

		assert.Nil(t, err)
		assert.Equal(t, enum.JourneyInProgress, state)
	})

	t.Run("when journey is valid, start it", func(t *testing.T) {
		target := createOrchestrator()

		car, state := target.Locate(JOURNEY_ID)

		assert.NotNil(t, car)
		assert.Equal(t, state, enum.JourneyInProgress)
		assert.Contains(t, car.JourneyIds, JOURNEY_ID)
	})

	t.Run("when journey is duplicated returns error", func(t *testing.T) {
		target := createOrchestrator()

		err := target.RegisterJourney(*models.NewJourney(JOURNEY_ID, 1))

		assert.IsType(t, &errors.ErrorJourneyAlreadyExists{}, err)
	})

	t.Run("when there are no available cars enqueue the journey", func(t *testing.T) {
		target := createOrchestrator()
		target.RegisterJourney(*models.NewJourney(2, 1000))

		_, state := target.Locate(2)

		assert.Equal(t, enum.JourneyQueued, state)
	})
}

func TestOrchestratorRegisterCars(t *testing.T) {
	t.Run("when cars are valid returns no error", func(t *testing.T) {
		target := createOrchestrator()

		err := target.RegisterCars([]models.Car{
			*models.NewCar(1, 1),
			*models.NewCar(2, 2),
			*models.NewCar(3, 3),
		})

		assert.Nil(t, err)

		assert.Empty(t, target.storage.FindCar(1).JourneyIds)
		assert.Empty(t, target.storage.FindCar(2).JourneyIds)
		assert.Empty(t, target.storage.FindCar(3).JourneyIds)
	})

	t.Run("when cars are duplicated returns error", func(t *testing.T) {
		target := createOrchestrator()

		err := target.RegisterCars([]models.Car{
			*models.NewCar(1, 1),
			*models.NewCar(1, 2),
		})

		assert.IsType(t, &errors.ErrorDuplicatedCarId{}, err)
	})
}

func createOrchestrator() *orchestrator {
	orchestrator := NewOrchestrator()
	orchestrator.RegisterCars([]models.Car{
		*models.NewCar(CAR_ID, CAR_SEATS),
	})

	orchestrator.RegisterJourney(*models.NewJourney(JOURNEY_ID, JOURNEY_PEOPLE))

	return orchestrator
}
