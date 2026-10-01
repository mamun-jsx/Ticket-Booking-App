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
	// DTO রিকুয়েস্ট থেকে ডেটা নিয়ে নতুন User মডেল তৈরি করা হচ্ছে
	newUser := User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	// রিপোজিটরি মেথড কল করে ডাটাবেজে ইউজার সেভ করা হচ্ছে
	err := s.repo.CreateUser(&newUser)

	// যদি ডাটাবেজে ইউজার তৈরিতে কোনো এরর হয়, তবে তা রিটার্ন করা হচ্ছে
	if err != nil {
		return nil, err
	}

	// সফলভাবে সেভ হওয়ার পর ক্লায়েন্টের জন্য রেসপন্স DTO প্রস্তুত করা হচ্ছে
	response := dto.Response{
		ID:        newUser.ID,
		Name:      newUser.Name,
		Email:     newUser.Email,
		CreatedAt: newUser.CreatedAt.String(),
	}

	// রেসপন্সের পয়েন্টার রিটার্ন করা হচ্ছে
	return &response, nil
}
