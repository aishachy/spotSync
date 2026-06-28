package parkingzone

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB) {

	parkingZoneRepository := NewRepository(db)
	parkingZoneService := NewService(parkingZoneRepository)
	parkingZoneHandler := NewHandler(parkingZoneService)

	api := e.Group("/api/v1/zones")

	api.POST("", parkingZoneHandler.CreateZone)
	api.GET("", parkingZoneHandler.GetAllZones)
	api.GET("/:id", parkingZoneHandler.GetZoneByID)
}
