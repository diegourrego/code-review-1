package repository

import "app/internal"

// NewVehicleMap is a function that returns a new instance of VehicleMap
func NewVehicleMap(db map[int]internal.Vehicle) *VehicleMap {
	// default db
	defaultDb := make(map[int]internal.Vehicle)
	if db != nil {
		defaultDb = db
	}
	return &VehicleMap{db: defaultDb}
}

// VehicleMap is a struct that represents a vehicle repository
type VehicleMap struct {
	// db is a map of vehicles
	db map[int]internal.Vehicle
}

// FindAll is a method that returns a map of all vehicles
func (r *VehicleMap) FindAll() (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	for key, value := range r.db {
		v[key] = value
	}

	return
}

func (r *VehicleMap) FindByID(vehicleId int) (exist bool) {
	_, ok := r.db[vehicleId]
	if !ok {
		return false
	}
	return true

}

func (r *VehicleMap) Create(newVehicle internal.Vehicle) error {
	carExists := r.FindByID(newVehicle.Id)
	if carExists {
		return internal.ErrCarAlreadyExists
	}

	r.db[newVehicle.Id] = newVehicle

	return nil
}
