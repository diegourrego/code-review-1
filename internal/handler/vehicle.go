package handler

import (
	"app/internal"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
)

// VehicleJSON is a struct that represents a vehicle in JSON format
type VehicleJSON struct {
	ID              int     `json:"id"`
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	Registration    string  `json:"registration"`
	Color           string  `json:"color"`
	FabricationYear int     `json:"year"`
	Capacity        int     `json:"passengers"`
	MaxSpeed        float64 `json:"max_speed"`
	FuelType        string  `json:"fuel_type"`
	Transmission    string  `json:"transmission"`
	Weight          float64 `json:"weight"`
	Height          float64 `json:"height"`
	Length          float64 `json:"length"`
	Width           float64 `json:"width"`
}

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(sv internal.VehicleService) *VehicleDefault {
	return &VehicleDefault{sv: sv}
}

// VehicleDefault is a struct with methods that represent handlers for vehicles
type VehicleDefault struct {
	// sv is the service that will be used by the handler
	sv internal.VehicleService
}

// GetAll is a method that returns a handler for the route GET /vehicles
func (h *VehicleDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		// - get all vehicles
		v, err := h.sv.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *VehicleDefault) CreateVehicle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid body")
		}

		var bodyMap map[string]any
		if err := json.Unmarshal(bytes, &bodyMap); err != nil {
			response.Text(w, http.StatusBadRequest, "invalid body")
			return
		}

		if err := validateIfKeysExist(bodyMap, "id", "brand", "model", "registration", "year", "color", "max_speed",
			"fuel_type", "transmission", "passengers", "height", "width", "weight"); err != nil {
			response.Text(w, http.StatusBadRequest, "invalid body. Keys are missing")
			return
		}

		// Deserialization of the body
		var vehicle internal.Vehicle
		if err := json.Unmarshal(bytes, &vehicle); err != nil {
			response.Text(w, http.StatusBadRequest, "invalid body")
			return
		}

		// Error handling
		if err := h.sv.CreateVehicle(vehicle); err != nil {
			switch {
			case errors.Is(err, internal.ErrInvalidBody):
				response.Text(w, http.StatusBadRequest, err.Error())
			case errors.Is(err, internal.ErrCarAlreadyExists):
				response.Text(w, http.StatusConflict, err.Error())
			}
			return
		}

		// response
		data := VehicleJSON{
			ID:              vehicle.Id,
			Brand:           vehicle.Brand,
			Model:           vehicle.Model,
			Registration:    vehicle.Registration,
			Color:           vehicle.Color,
			FabricationYear: vehicle.FabricationYear,
			Capacity:        vehicle.Capacity,
			MaxSpeed:        vehicle.MaxSpeed,
			FuelType:        vehicle.FuelType,
			Transmission:    vehicle.Transmission,
			Weight:          vehicle.Weight,
			Height:          vehicle.Height,
			Length:          vehicle.Length,
			Width:           vehicle.Width,
		}

		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "Vehicle created successfully",
			"data":    data,
		})

	}
}

func (h *VehicleDefault) FindByColorAndYear() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtain color and year
		color := chi.URLParam(r, "color")
		yearStr := chi.URLParam(r, "year")
		// Validate year
		year, err := strconv.Atoi(yearStr)
		if err != nil {
			response.Text(w, http.StatusBadRequest, "Invalid year. Year must be a numeric value")
			return
		}

		vehiclesFounded, err := h.sv.FindByColorAndYear(color, year)
		if err != nil {
			response.Text(w, http.StatusNotFound, err.Error())
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Cars list obtained",
			"data":    vehiclesFounded,
		})
	}
}

func (h *VehicleDefault) FindByBrandAndYearRate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		brand := chi.URLParam(r, "brand")
		initialYearStr := chi.URLParam(r, "start_year")
		finalYearStr := chi.URLParam(r, "end_year")
		initialYear, err := strconv.Atoi(initialYearStr)
		if err != nil {
			response.Text(w, http.StatusBadRequest, "Invalid year. Year must be a numeric value")
			return
		}

		finalYear, err := strconv.Atoi(finalYearStr)
		if err != nil {
			response.Text(w, http.StatusBadRequest, "Invalid year. Year must be a numeric value")
			return
		}

		vehiclesFounded, err := h.sv.FindBetweenBrandAndYearRate(brand, initialYear, finalYear)
		if err != nil {
			response.Text(w, http.StatusNotFound, err.Error())
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "cars list founded",
			"data":    vehiclesFounded,
		})

	}
}

func (h *VehicleDefault) FindVelocityAverageByBrand() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		brand := chi.URLParam(r, "brand")
		brandVelocityAverage, err := h.sv.FindVelocityAverageByBrand(brand)
		if err != nil {
			response.Text(w, http.StatusNotFound, err.Error())
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message":           "Speed average found",
			"brandSpeedAverage": brandVelocityAverage,
		})
	}
}

func (h *VehicleDefault) CreateVehicles() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid body")
		}

		var vehicles []internal.Vehicle
		if err := json.Unmarshal(bytes, &vehicles); err != nil {
			response.Text(w, http.StatusBadRequest, "invalid body")
			return
		}

		// Validación no funciona :(
		//if err := validateAllVehiclesKeys(vehicles); err != nil {
		//	response.Text(w, http.StatusBadRequest, "invalid body. Keys are missing")
		//	return
		//}

		if err := h.sv.CreateVehicules(vehicles); err != nil {
			switch {
			case errors.Is(err, internal.ErrInvalidBody):
				response.Text(w, http.StatusBadRequest, err.Error())
			case errors.Is(err, internal.ErrCarAlreadyExists):
				response.Text(w, http.StatusConflict, err.Error())
			}
			return
		}

		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "Vehicules created successfully",
			"data":    vehicles,
		})

	}
}

func (h *VehicleDefault) UpdateMaxSpeed() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid id. Id must be an int number.")
			return
		}

		var bodyMap map[string]any

		if err := json.NewDecoder(r.Body).Decode(&bodyMap); err != nil {
			response.Text(w, http.StatusBadRequest, "invalid body")
			return
		}

		// Con esto me aseguro que venga en la request max_speed
		if err := validateIfKeysExist(bodyMap, "max_speed"); err != nil {
			response.Text(w, http.StatusBadRequest, "invalid request. max_speed must be provided.")
			return
		}

		newMaxSpeed, ok := bodyMap["max_speed"].(float64)
		if !ok {
			response.Text(w, http.StatusBadRequest, "Invalid max_speed value")
			return
		}

		vehicleUpdated, err := h.sv.UpdateMaxSpeed(id, newMaxSpeed)
		if err != nil {

			if errors.Is(err, internal.ErrVehicleNotFounded) {
				response.Text(w, http.StatusNotFound, err.Error())
				return
			}

			response.Text(w, http.StatusBadRequest, err.Error())
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message":        "max_speed updated successfully",
			"vehicleUpdated": vehicleUpdated,
		})

	}
}

func (h *VehicleDefault) FindVehiclesByFuelType() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fuelType := chi.URLParam(r, "type")

		vehiclesFounded, err := h.sv.FindVehiclesByFuelType(fuelType)
		if err != nil {
			response.Text(w, http.StatusNotFound, err.Error())
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Vehicles found!",
			"data":    vehiclesFounded,
		})
	}
}

func (h *VehicleDefault) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			response.Text(w, http.StatusBadRequest, "Invalid ID format. ID must be an int number.")
			return
		}

		if err := h.sv.Delete(id); err != nil {
			response.Text(w, http.StatusNotFound, err.Error())
			return
		}

		response.JSON(w, http.StatusNoContent, map[string]any{
			"message": "vehicle deleted successfully",
		})
	}
}

func (h *VehicleDefault) FindVehiculesByTransmissionType() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		transmissionType := chi.URLParam(r, "type")

		vehiclesFound, err := h.sv.FindVehiculesByTransmissionType(transmissionType)
		if err != nil {
			response.Text(w, http.StatusNotFound, err.Error())
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "vehicles found!",
			"data":    vehiclesFound,
		})
	}
}

func (h *VehicleDefault) UpdateFuelType() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			response.Text(w, http.StatusBadRequest, "Invalid ID format. ID must be an int number.")
			return
		}

		var bodyMap map[string]any
		if err := json.NewDecoder(r.Body).Decode(&bodyMap); err != nil {
			response.Text(w, http.StatusBadRequest, "Invalid body format")
			return
		}

		newFuelType, ok := bodyMap["fuel_type"].(string)
		if !ok {
			response.Text(w, http.StatusBadRequest, "fuel type must be provided.")
			return
		}

		vehicleUpdated, err := h.sv.UpdateFuelType(id, newFuelType)
		if err != nil {
			response.Text(w, http.StatusNotFound, err.Error())
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message":        "fuel type updated successfully in vehicle",
			"vehicleUpdated": vehicleUpdated,
		})

	}
}

func validateIfKeysExist(data map[string]any, keys ...string) error {
	for _, key := range keys {
		if _, ok := data[key]; !ok {
			return internal.ErrInvalidBody
		}
	}
	return nil
}
