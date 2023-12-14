package api

import (
	"net/http"
	"strconv"

	"github.com/imarcelolz/transportation-challenge/services/carpool"
	"github.com/imarcelolz/transportation-challenge/services/carpool/enum"
	"github.com/imarcelolz/transportation-challenge/services/carpool/models"
)

type poolController struct {
	orchestrator carpool.Orchestrator
}

func NewPoolController() *poolController {
	return &poolController{
		orchestrator: carpool.NewOrchestrator(),
	}
}

func (c *poolController) cars(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	cars := decodeJSON[[]models.Car](r.Body)

	c.orchestrator = carpool.NewOrchestrator()

	if err := c.orchestrator.RegisterCars(cars); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *poolController) dropoff(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	id, ok := getFormId(w, r)
	if !ok {
		return
	}

	if err := c.orchestrator.Dropoff(id); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *poolController) journey(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	request := decodeJSON[models.Journey](r.Body)

	if err := c.orchestrator.RegisterJourney(request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func (c *poolController) locate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	id, ok := getFormId(w, r)
	if !ok {
		return
	}

	car, state := c.orchestrator.Locate(id)
	if state == enum.JourneyQueued {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if state == enum.JourneyUndefined {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	respondToJson[*models.Car](w, car)
}

func getFormId(w http.ResponseWriter, r *http.Request) (int, bool) {
	r.ParseForm()
	raw := r.FormValue("ID")

	if raw == "" {
		raw = r.FormValue("id")
	}

	id, err := strconv.Atoi(raw)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return 0, false
	}

	return id, true
}
