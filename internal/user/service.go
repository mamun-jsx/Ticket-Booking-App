package user

import "github.com/mamun-jsx/Ticket-Booking-App/internal/user/dto"

// Service handles business logic related to users.
type Service struct {
	repo UserRepository
}

// NewService creates and returns a new instance of Service as a pointer.
func NewService(repo UserRepository) *Service {
	return &Service{
		repo: repo,
	}
}

// CreateUser handles user creation logic, saves the user to repository, and returns a response DTO.
func (s *Service) CreateUser(req *dto.CreateRequest) (*dto.Response, error) {
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
