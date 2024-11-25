package usecase

import (
	"fmt"
	"test-mnc/repository"
	"test-mnc/shared/model"
	"test-mnc/shared/service"
)

type blacklistUsecase struct {
	repo       repository.BlacklistRepository
	jwtService service.JwtService
}

// AddToken implements BlacklistUsecase.
func (b *blacklistUsecase) AddTokenToBlacklist(token string) error {
	return b.repo.AddTokenToBlacklist(token)
}

// IsBlacklisted implements BlacklistUsecase.
func (b *blacklistUsecase) ValidateAndProcessToken(token string) (*model.MyCustomClaims, error) {
	// Check if token is blacklisted
	isBlacklisted, err := b.repo.IsBlacklisted(token)
	if err != nil {
		return nil, fmt.Errorf("Failed to check blacklist: %v", err)
	}
	if isBlacklisted {
		return nil, fmt.Errorf("Token is blacklisted")
	}

	claims, err := b.jwtService.ValidateToken(token)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

type BlacklistUsecase interface {
	AddTokenToBlacklist(token string) error
	ValidateAndProcessToken(token string) (*model.MyCustomClaims, error)
}

func NewBlacklistUsecase(repo repository.BlacklistRepository, jwtService service.JwtService) BlacklistUsecase {
	return &blacklistUsecase{repo: repo, jwtService: jwtService}
}
