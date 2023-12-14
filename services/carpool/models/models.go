package models

import (
	"github.com/imarcelolz/transportation-challenge/services/carpool/tools"
)

type Car struct {
	Id             int   `json:"id"`
	Seats          int   `json:"seats"`
	AvailableSeats int   `json:"-"`
	JourneyIds     []int `json:"-"`
}

func NewCar(id, seats int) *Car {
	return &Car{
		Id:             id,
		Seats:          seats,
		AvailableSeats: seats,
		JourneyIds:     []int{},
	}
}

func (c *Car) AddJourney(journey *Journey) {
	c.JourneyIds = append(c.JourneyIds, journey.Id)
	c.AvailableSeats -= journey.People
	journey.Car = c
}

func (c *Car) RemoveJourney(journey *Journey) {
	c.JourneyIds = tools.ArrayRemove(c.JourneyIds, journey.Id)
	c.AvailableSeats += journey.People
	journey.Car = nil
}

type Journey struct {
	Id     int  `json:"id"`
	People int  `json:"people"`
	Car    *Car `json:"-"`
}

func NewJourney(id, people int) *Journey {
	return &Journey{
		Id:     id,
		People: people,
	}
}
