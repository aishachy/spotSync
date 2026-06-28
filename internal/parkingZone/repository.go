package parkingzone

import (
	"errors"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	CreateZone(zone *ParkingZone) error
	GetAllZones() ([]ParkingZone, error)
	GetZoneByID(id uint) (*ParkingZone, error)
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r repository) CreateZone(zone *ParkingZone) error {
	return r.db.Create(zone).Error
}

func (r repository) GetAllZones() ([]ParkingZone, error) {
	var zones []ParkingZone

	err := r.db.Find(&zones).Error
	if err != nil {
		return nil, err
	}

	return zones, nil
}

func (r repository) GetZoneByID(id uint) (*ParkingZone, error) {
	var zone ParkingZone

	result := r.db.First(&zone, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &zone, nil
}
