package service

import (
	"fmt"
	"test-mnc/config"
	"test-mnc/entity"
	"test-mnc/entity/dto"
	"test-mnc/shared/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtService struct {
	cfg config.TokenConfig
}

// GenerateToken implements JwtService.
func (j *jwtService) GenerateToken(author entity.Customer) (dto.AuthResponseDto, error) {
	claims := model.MyCustomClaims{
		Name:  author.Name,
		Email: author.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.cfg.IssuerName,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.cfg.JwtExpiresTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(j.cfg.JwtSigningMethod, claims)
	ss, err := token.SignedString(j.cfg.JwtSignatureKy)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}

	return dto.AuthResponseDto{Token: ss}, nil
}

// ValidateToken implements JwtService.
func (j *jwtService) ValidateToken(tokenString string) (*model.MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.MyCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return j.cfg.JwtSignatureKy, nil
	})
	if err != nil {
		return nil, fmt.Errorf("Unauthorized")
	}

	if !token.Valid {
		return nil, fmt.Errorf("Unauthorized")
	}

	claims, ok := token.Claims.(*model.MyCustomClaims)
	if !ok {
		return nil, fmt.Errorf("Invalid token")
	}

	return claims, nil
}

type JwtService interface {
	GenerateToken(author entity.Customer) (dto.AuthResponseDto, error)
	ValidateToken(tokenString string) (*model.MyCustomClaims, error)
}

func NewJwtService(cfg config.TokenConfig) JwtService {
	return &jwtService{cfg: cfg}
}
