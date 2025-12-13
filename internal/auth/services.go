package auth

import (
	"errors"

	"github.com/google/uuid"
	"github.com/lukiriskigumilar/resepify-be/internal/users"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	RegisterService(input RegisterRequestDTO) (*users.User, error)
}

type authService struct {
	repo users.UserRepository
}

func NewAuthService(repo users.UserRepository) AuthService {
	return &authService{repo}
}

func (s *authService) RegisterService(input RegisterRequestDTO) (*users.User, error) {

	// check if email already exists
	existing, _ := s.repo.FindByEmail(input.Email)
	if existing.ID != uuid.Nil {
		return nil, errors.New("email already in use")
	}

	//hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// create user
	newUser := &users.User{
		ID:       uuid.New(),
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashed),
	}

	err = s.repo.Create(newUser)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}
