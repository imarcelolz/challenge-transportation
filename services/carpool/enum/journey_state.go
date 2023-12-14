package enum

type JourneyState int

const (
	JourneyUndefined JourneyState = iota
	JourneyQueued
	JourneyInProgress
)
