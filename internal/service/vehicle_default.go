package service

import "app/internal"

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(rp internal.VehicleRepository) *VehicleDefault {
	return &VehicleDefault{rp: rp}
}

// VehicleDefault is a struct that represents the default service for vehicles
type VehicleDefault struct {
	// rp is the repository that will be used by the service
	rp internal.VehicleRepository
}

// FindAll is a method that returns a map of all vehicles
func (s *VehicleDefault) FindAll() (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindAll()
	return
}

func (s *VehicleDefault) CreateVehicle(newVehicle internal.Vehicle) error {
	if err := s.rp.CreateVehicle(newVehicle); err != nil {
		return err
	}

	return nil
}

func (s *VehicleDefault) FindByColorAndYear(color string, year int) (map[int]internal.Vehicle, error) {
	vehicles, err := s.rp.FindByColorAndYear(color, year)
	if err != nil {
		return nil, err
	}
	return vehicles, nil
}

func (s *VehicleDefault) FindBetweenBrandAndYearRate(brand string, initialYear int, finalYear int) (map[int]internal.Vehicle, error) {
	vehicles, err := s.rp.FindBetweenBrandAndYearRate(brand, initialYear, finalYear)
	if err != nil {
		return nil, err
	}
	return vehicles, nil
}

func (s *VehicleDefault) FindVelocityAverageByBrand(brand string) (float64, error) {
	brandVelocityAverage, err := s.rp.FindVelocityAverageByBrand(brand)
	if err != nil {
		return 0, err
	}
	return brandVelocityAverage, nil
}

func (s *VehicleDefault) CreateVehicules(newVehicles []internal.Vehicle) error {
	if err := s.rp.CreateVehicules(newVehicles); err != nil {
		return err
	}
	return nil
}
