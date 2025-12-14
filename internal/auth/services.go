package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/lukiriskigumilar/resepify-be/internal/users"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	RegisterService(input RegisterRequestDTO) (*users.User, error)
	LoginService(input LoginRequestDTO) (*LoginResponseDTO, error)
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

func (s *authService) LoginService(input LoginRequestDTO) (*LoginResponseDTO, error) {

	user, err := s.repo.FindByEmail(input.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	//compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	ProcessesToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	dsn := os.Getenv("JWT_SECRET")

	token, err := ProcessesToken.SignedString([]byte(dsn))
	if err != nil {
		return nil, err
	}

	responseData := &LoginResponseDTO{
		Name:        user.Name,
		Email:       user.Email,
		Token:       token,
		TokenExpiry: time.Now().Add(time.Hour * 24),
	}

	return responseData, nil
}
