package auth

import (
	"audio_converter/models"
	"audio_converter/utils"
	"encoding/hex"
	"errors"
)

var (
	ErrInvalidCredentials = errors.New("неверный email или пароль")
	ErrTokenCreation      = errors.New("ошибка при создании токена")
)

type AuthUseCase struct {
	userRepo   models.UserRepository
	accessRepo models.AccessRepository
	tokenGen   models.TokenGenerator
}

func NewAuthUseCase(
	userRepo models.UserRepository,
	accessRepo models.AccessRepository,
	tokenGen models.TokenGenerator,
) *AuthUseCase {
	return &AuthUseCase{
		userRepo:   userRepo,
		accessRepo: accessRepo,
		tokenGen:   tokenGen,
	}
}

func (uc *AuthUseCase) Login(email, password string) (*models.User, *models.Token, error) {
	user, err := uc.userRepo.FindByEmail(email)
	if err != nil {
		return nil, nil, ErrInvalidCredentials
	}

	passHash, err := hex.DecodeString(user.Password)
	if err != nil {
		return nil, nil, err
	}

	if !utils.CheckPass(passHash, password) {
		return nil, nil, ErrInvalidCredentials
	}

	token := uc.tokenGen.GenerateToken()
	userAccess := &models.Access{
		UserID: user.ID,
		Token:  token,
	}

	if err := uc.accessRepo.CreateAccess(userAccess); err != nil {
		return nil, nil, ErrTokenCreation
	}

	return user, token, nil
}
