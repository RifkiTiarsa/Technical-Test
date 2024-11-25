package usecase

import (
	"test-mnc/entity"
	"test-mnc/entity/dto"
	"test-mnc/shared/service"
)

type authUseCase struct {
	uc         CustomerUsecase
	jwtService service.JwtService
}

type AuthUsecase interface {
	Login(payload dto.AuthRequestDto) (dto.AuthResponseDto, error)
	Register(payload entity.Customer) (entity.Customer, error)
}

// Login implements AuthUsecase.
func (a *authUseCase) Login(payload dto.AuthRequestDto) (dto.AuthResponseDto, error) {
	customer, err := a.uc.FindCustomerByEmailPassword(payload.Email, payload.Password)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}

	token, err := a.jwtService.GenerateToken(customer)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}

	return dto.AuthResponseDto{Token: token.Token}, nil
}

// Register implements AuthUsecase.
func (a *authUseCase) Register(payload entity.Customer) (entity.Customer, error) {
	return a.uc.RegisterCustomer(payload)
}

// NewAuthUsecase returns new instance of AuthUsecase.
func NewAuthUsecase(uc CustomerUsecase, jwtService service.JwtService) AuthUsecase {
	return &authUseCase{uc: uc, jwtService: jwtService}
}
