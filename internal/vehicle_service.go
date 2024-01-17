package internal

import "errors"

var (
	ErrInvalidBody = errors.New("invalid request body. Please check it and try again")
)

// VehicleService is an interface that represents a vehicle service
type VehicleService interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	// CreateVehicle creates a new vehicle in memory - requirement 1
	CreateVehicle(newVehicle Vehicle) error
	// FindByColorAndYear filters cars according year and color - requirement 2
	FindByColorAndYear(color string, year int) (map[int]Vehicle, error)
	// FindBetweenBrandAndYearRate filters cars according a specific brand and year rate - requirement 3
	FindBetweenBrandAndYearRate(brand string, initialYear int, finalYear int) (map[int]Vehicle, error)
	// FindVelocityAverageByBrand finds an average of a specific brand - requirement 4
	FindVelocityAverageByBrand(brand string) (float64, error)
	// CreateVehicules creates many vehicules - requirement 5
	CreateVehicules(newVehicles []Vehicle) error
	// UpdateMaxSpeed update only vehicle max_speed - requirement 6
	UpdateMaxSpeed(vehicleID int, newMaxSpeed float64) (Vehicle, error)
	// FindVehiclesByFuelType finds vehicles by fuel type - requirement 7
	FindVehiclesByFuelType(fuelType string) (v map[int]Vehicle, err error)
	// Delete deletes a vehicle - requirement 8
	Delete(vehicleID int) error
}
