package server

import (
	"fmt"
	"net/http"

	"spotSync/internal/config"
	"spotSync/internal/parkingzone"
	"spotSync/internal/user"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.ErrBadRequest.Wrap(err)
	}
	return nil
}

func Start(db *gorm.DB, cfg *config.Config) {

	// migrate tables
	db.AutoMigrate(
		&user.User{},
		&parkingzone.ParkingZone{},
	)

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	e.GET("/", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Hello, World!",
		})
	})

	// register routes
	user.RegisterRoutes(e, db)
	parkingzone.RegisterRoutes(e, db)

	port := fmt.Sprintf(":%s", cfg.Port)
	e.Start(port)
}
