package internal

import "errors"

var (
	ErrInvalidBody = errors.New("invalid request body. Please check it and try again")
)

// VehicleService is an interface that represents a vehicle service
type VehicleService interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	// Create creates a new vehicle in memory - requirement 1
	Create(newVehicle Vehicle) error
	// FindByColorAndYear filters cars according year and color - requirement 2
	FindByColorAndYear(color string, year int) (map[int]Vehicle, error)
}
