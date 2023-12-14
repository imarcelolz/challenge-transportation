package carpool

import (
	"math/rand"
	"testing"

	"github.com/imarcelolz/transportation-challenge/services/carpool/models"
)

func BenchmarkDropoff(b *testing.B) {
	service, journeys := newPopulatedService(b.N)

	for i := 0; i < b.N; i++ {
		service.Dropoff(journeys[i].Id)
	}
}

func BenchmarkLocate(b *testing.B) {
	service, journeys := newPopulatedService(b.N)

	for i := 0; i < b.N; i++ {
		journey := journeys[rand.Intn(len(journeys))]
		service.Locate(journey.Id)
	}
}

func BenchmarkRegisterCars(b *testing.B) {
	b.Run("Register Cars Benchmark", func(b *testing.B) {
		service := NewOrchestrator()

		for i := 0; i < b.N; i++ {
			service.RegisterCars(buildCars(b.N))
		}
	})
}

func BenchmarkRegisterJourneys(b *testing.B) {
	service, existingJourneys := newPopulatedService(b.N)
	journeys := buildJourneys(b.N, len(existingJourneys))

	for i := 0; i < b.N; i++ {
		journey := journeys[i]
		service.RegisterJourney(journey)
	}
}

func newPopulatedService(count int) (Orchestrator, []models.Journey) {
	cars := buildCars(count)
	journeys := buildJourneys(count, 0)

	service := NewOrchestrator()
	service.RegisterCars(cars)

	for _, journey := range journeys {
		service.RegisterJourney(journey)
	}

	return service, journeys
}

func buildCars(count int) []models.Car {
	cars := make([]models.Car, count)

	for i := 0; i < count; i++ {
		cars[i] = *models.NewCar(i+1, rand.Intn(20))
	}

	return cars
}

func buildJourneys(count int, startIndex int) []models.Journey {
	journeys := make([]models.Journey, count)

	for i := 0; i < count; i++ {
		journeys[i] = *models.NewJourney(i+1+startIndex, rand.Intn(30))
	}

	return journeys
}
