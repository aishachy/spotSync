package reservation

import (
	"spotSync/internal/parkingzone"
	"spotSync/internal/user"

	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model

	UserID uint      `json:"user_id" gorm:"not null"`
	User   user.User `gorm:"foreignKey:UserID"`

	ZoneID uint                    `json:"zone_id" gorm:"not null"`
	Zone   parkingzone.ParkingZone `gorm:"foreignKey:ZoneID"`

	LicensePlate string `json:"license_plate" gorm:"type:varchar(15);not null"`

	Status string `json:"status" gorm:"type:varchar(20);default:active"`
}
