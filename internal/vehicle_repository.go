package internal

import "errors"

var (
	ErrCarAlreadyExists  = errors.New("vehicle identifier already exists")
	ErrVehicleNotFounded = errors.New("none vehicles found according criteria")
)

// VehicleRepository is an interface that represents a vehicle repository
type VehicleRepository interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	FindByID(vehicleId int) (exist bool)
	// Create creates a new vehicle in memory - requirement 1
	Create(newVehicle Vehicle) error
	// FindByColorAndYear filters cars according year and color - requirement 2
	FindByColorAndYear(color string, year int) (map[int]Vehicle, error)
	// FindBetweenBrandAndYearRate filters cars according a specific brand and year rate - requirement 3
	FindBetweenBrandAndYearRate(brand string, initialYear int, finalYear int) (map[int]Vehicle, error)
}
