package user

import (
	"fmt"
	"spotSync/internal/auth"
	"spotSync/internal/user/dto"
)

var ErrInvalidCredentials = fmt.Errorf("Invalid email or password")

type service struct {
	repo       Repository
	jwtService auth.JWTService
}

func NewService(repo Repository, jwtService auth.JWTService) *service {
	return &service{
		repo:       repo,
		jwtService: jwtService,
	}
}

func (s *service) CreateUser(req dto.CreateRequest) (*dto.UserResponse, error) {

	user := User{
		Name:  req.Name,
		Email: req.Email,
		Role:  req.Role,
	}

	if user.Role == "" {
		user.Role = "driver"
	}
	err := user.hashedPassword(req.Password)
	if err != nil {
		return nil, err
	}

	err = s.repo.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	response := dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}

	return &response, nil
}

func (s *service) LoginUser(req dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	err = user.checkPassword(req.Password)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	token, err := s.jwtService.GenerateToken(user.ID, user.Email, user.Name)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate token: %w", err)

	}

	response := dto.LoginResponse{
		Token: token,
		User: dto.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role,
		},
	}

	return &response, nil

}
