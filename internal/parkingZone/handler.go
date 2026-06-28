package parkingzone

import (
	"net/http"
	"strconv"
	"strings"

	httpresponse "spotSync/internal/httpResponse"
	"spotSync/internal/parkingzone/dto"

	"github.com/labstack/echo/v5"
)

type handler struct {
	service *service
}

func NewHandler(service *service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) CreateZone(c *echo.Context) error {
	var req dto.CreateRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid request payload",
			Details: err.Error(),
		})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Validation failed",
			Details: err.Error(),
		})
	}

	response, err := h.service.CreateZone(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpresponse.Error{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create parking zone",
			Details: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, httpresponse.Success{
		Success: true,
		Message: "Parking zone created successfully",
		Data:    response,
	})
}

func (h *handler) GetAllZones(c *echo.Context) error {

	response, err := h.service.GetAllZones()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpresponse.Error{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve parking zones",
			Details: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, httpresponse.Success{
		Success: true,
		Message: "Parking zones retrieved successfully",
		Data:    response,
	})
}

func (h *handler) GetZoneByID(c *echo.Context) error {

	// Example URL: /api/v1/zones/2
	path := c.Request().URL.Path

	parts := strings.Split(path, "/")
	if len(parts) < 1 {
		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid parking zone ID",
		})
	}

	idStr := parts[len(parts)-1]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpresponse.Error{
			Code:    http.StatusBadRequest,
			Message: "Invalid parking zone ID",
			Details: err.Error(),
		})
	}

	response, err := h.service.GetZoneByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpresponse.Error{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve parking zone",
			Details: err.Error(),
		})
	}

	if response == nil {
		return c.JSON(http.StatusNotFound, httpresponse.Error{
			Code:    http.StatusNotFound,
			Message: "Parking zone not found",
			Details: "No parking zone found with this ID",
		})
	}

	return c.JSON(http.StatusOK, httpresponse.Success{
		Success: true,
		Message: "Parking zone retrieved successfully",
		Data:    response,
	})
}
