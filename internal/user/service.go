package user

import "github.com/mamun-jsx/Ticket-Booking-App/internal/user/dto"

// UserService defines the business logic interface for user operations.
type UserService interface {
	CreateUser(req *dto.CreateRequest) (*dto.Response, error)
}

// service implements UserService interface.
type service struct {
	repo UserRepository
}

// NewService creates and returns a new UserService implementation.
func NewService(repo UserRepository) UserService {
	return &service{
		repo: repo,
	}
}

// CreateUser handles user creation logic, saves the user to repository, and returns a response DTO.
func (s *service) CreateUser(req *dto.CreateRequest) (*dto.Response, error) {
	// 1. Map incoming DTO request data to the domain User entity
	newUser := User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	// 2. Persist the new user entity in the database via the repository layer
	err := s.repo.CreateUser(&newUser)

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
