package dto

type CreateRequest struct {
	UserID uint `json:"user_id" validate:"required"`

	ZoneID uint `json:"zone_id" validate:"required"`

	LicensePlate string `json:"license_plate" validate:"required,max=15"`
}
