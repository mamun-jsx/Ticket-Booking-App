package user

import (
	"errors"
	"fmt"

	"github.com/mamun-jsx/Ticket-Booking-App/internal/auth"
	"github.com/mamun-jsx/Ticket-Booking-App/internal/user/dto"
)

var ErrorInvalidEmailPassword = errors.New("Invalid Access ID password Not Match")

// UserService defines the business logic interface for user operations.
type UserService interface {
	CreateUser(req *dto.CreateRequest) (*dto.Response, error)
	LoginUser(req *dto.LoginRequest) (*dto.Response, error)
}

// service implements UserService interface.
type service struct {
	repo       UserRepository
	jwtService auth.JWTService
}

// NewService creates and returns a new UserService implementation.
func NewService(repo UserRepository, jwtService auth.JWTService) UserService {
	return &service{repo, jwtService}
}

// CreateUser handles user creation logic, saves the user to repository, and returns a response DTO.
func (s *service) CreateUser(req *dto.CreateRequest) (*dto.Response, error) {
	// 1. Map incoming DTO request data to the domain User entity
	newUser := User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	err := newUser.hashPassword(req.Password) // convert raw password to hashed
	if err != nil {
		return nil, err
	}

	// 2. Persist the new user entity in the database via the repository layer
	err = s.repo.CreateUser(&newUser)
	// 3. Return an error if database creation fails (e.g. duplicate key or DB connection issues)
	if err != nil {
		return nil, err
	}

	// 4. Construct the outgoing response DTO, omitting sensitive fields like password
	response := dto.Response{
		ID:        newUser.ID,
		Name:      newUser.Name,
		Email:     newUser.Email,
		CreatedAt: newUser.CreatedAt.String(),
	}

	// 5. Return a pointer to the populated response DTO
	return &response, nil
}

// login user

func (s *service) LoginUser(req *dto.LoginRequest) (*dto.Response, error) {
	// query into db check email is exist or not into repo
	user, err := s.repo.GetUserByEmail(req.Email)

	// if user not found
	if err != nil {
		return nil, err
	}

	// if user is nil
	if user == nil {
		return nil, ErrorInvalidEmailPassword
	}

	// compare password

	err = user.checkPassword(req.Password)
	if err != nil {
		return nil, ErrorInvalidEmailPassword
	}
	// generate JWT token
	token, err := s.jwtService.GenerateToken(user.ID, user.Email, user.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token %w", err)
	}

	// response
	response := dto.Response{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Token:     token,
		CreatedAt: user.CreatedAt.String(),
	}
	return &response, nil

}
