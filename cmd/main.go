package main

import (
	"net/http"
	"spotSync/internal/user"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" validate:"required" gorm:"type:varchar(100); not null"`
	Email    string `json:"email" validate:"required,email" gorm:"type:varchar(255); uniqueIndex:not null"`
	Password string `json:"password" validate:"required,min=6" gorm:"type:varchar(100); not null"`
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.ErrBadRequest.Wrap(err)
	}
	return nil
}

func main() {

	dsn := "host=localhost user=postgres password=postgres dbname=spotSync port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})

	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	e.GET("/", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Hello, World!"})
	})

	e.Validator = &CustomValidator{validator: validator.New()}

	user.RegisterRoutes(e, db)

	if err := e.Start(":5000"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
