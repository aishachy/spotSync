package parkingzone

import "spotSync/internal/parkingzone/dto"

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateZone(req dto.CreateRequest) (*dto.Response, error) {
	zone := ParkingZone{
		Name:          req.Name,
		Type:          req.Type,
		TotalCapacity: req.TotalCapacity,
		PricePerHour:  req.PricePerHour,
	}

	err := s.repo.CreateZone(&zone)
	if err != nil {
		return nil, err
	}

	response := dto.Response{
		ID:            zone.ID,
		Name:          zone.Name,
		Type:          zone.Type,
		TotalCapacity: zone.TotalCapacity,
		PricePerHour:  zone.PricePerHour,
		CreatedAt:     zone.CreatedAt,
		UpdatedAt:     zone.UpdatedAt,
	}

	return &response, nil
}

func (s *service) GetAllZones() ([]dto.Response, error) {
	zones, err := s.repo.GetAllZones()
	if err != nil {
		return nil, err
	}

	response := make([]dto.Response, 0, len(zones))

	for _, zone := range zones {
		response = append(response, dto.Response{
			ID:            zone.ID,
			Name:          zone.Name,
			Type:          zone.Type,
			TotalCapacity: zone.TotalCapacity,
			PricePerHour:  zone.PricePerHour,
			CreatedAt:     zone.CreatedAt,
			UpdatedAt:     zone.UpdatedAt,
		})
	}

	return response, nil
}

func (s *service) GetZoneByID(id uint) (*dto.Response, error) {
	zone, err := s.repo.GetZoneByID(id)
	if err != nil {
		return nil, err
	}

	if zone == nil {
		return nil, nil
	}

	response := dto.Response{
		ID:            zone.ID,
		Name:          zone.Name,
		Type:          zone.Type,
		TotalCapacity: zone.TotalCapacity,
		PricePerHour:  zone.PricePerHour,
		CreatedAt:     zone.CreatedAt,
		UpdatedAt:     zone.UpdatedAt,
	}

	return &response, nil
}
