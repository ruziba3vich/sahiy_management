package auth

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ruziba3vich/sahiy_management/internal/domain/privilege"
	"github.com/ruziba3vich/sahiy_management/internal/domain/user"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidToken       = errors.New("invalid token")
	ErrTokenExpired       = errors.New("token expired")
)

type Claims struct {
	UserID   int64  `json:"user_id"`
	Phone    string `json:"phone"`
	FullName string `json:"full_name"`
	Role     int    `json:"role"`
	jwt.RegisteredClaims
}

type Service struct {
	userRepo      user.Repository
	privilegeRepo privilege.Repository
	jwtSecret     []byte
	jwtExpiry     time.Duration
}

func NewService(userRepo user.Repository, privilegeRepo privilege.Repository, jwtSecret string, jwtExpiryHours int) *Service {
	return &Service{
		userRepo:      userRepo,
		privilegeRepo: privilegeRepo,
		jwtSecret:     []byte(jwtSecret),
		jwtExpiry:     time.Duration(jwtExpiryHours) * time.Hour,
	}
}

func (s *Service) Login(ctx context.Context, phone, password string) (string, *user.User, error) {
	u, err := s.userRepo.GetUserByPhone(ctx, phone)
	if err != nil {
		return "", nil, ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)); err != nil {
		return "", nil, ErrInvalidCredentials
	}

	token, err := s.generateToken(u)
	if err != nil {
		return "", nil, err
	}

	return token, u, nil
}

func (s *Service) generateToken(u *user.User) (string, error) {
	claims := &Claims{
		UserID:   u.ID,
		Phone:    u.Phone,
		FullName: u.FullName,
		Role:     u.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.jwtExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.jwtSecret)
}

func (s *Service) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return s.jwtSecret, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrTokenExpired
		}
		return nil, ErrInvalidToken
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}

	return claims, nil
}

func (s *Service) RefreshToken(ctx context.Context, tokenString string) (string, error) {
	claims, err := s.ValidateToken(tokenString)
	if err != nil && !errors.Is(err, ErrTokenExpired) {
		return "", err
	}

	u, err := s.userRepo.GetUserByID(ctx, claims.UserID)
	if err != nil {
		return "", ErrInvalidToken
	}

	return s.generateToken(u)
}

func (s *Service) HasPrivilege(ctx context.Context, userID int64, role int, resource, action string) (bool, error) {
	if role == privilege.RoleSuperAdmin {
		return true, nil
	}

	return s.privilegeRepo.HasPrivilege(ctx, userID, resource, action)
}

func (s *Service) GetUserByID(ctx context.Context, id int64) (*user.User, error) {
	return s.userRepo.GetUserByID(ctx, id)
}
