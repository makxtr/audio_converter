package auth

import (
	"audio_converter/models"
	"audio_converter/repository"
	"audio_converter/utils"
	"encoding/hex"
	"testing"
)

func TestAuthUseCase_Success(t *testing.T) {
	salt := utils.GetSalt()

	hashedPassword := utils.HashPass(salt, "password")
	user := &models.User{
		ID:       1,
		Email:    "test@example.com",
		Password: hex.EncodeToString(hashedPassword),
	}
	repo := &repository.MockUserRepository{User: user}

	token := &models.Token{}
	access := &models.Access{
		UserID: user.ID,
		Token:  token,
	}
	accessRepo := &repository.MockAccessRepository{Access: access}
	tokenGenerator := &repository.MockTokenGenerator{Token: token}

	authUC := NewAuthUseCase(repo, accessRepo, tokenGenerator)

	lUser, lToken, _ := authUC.Login("test@example.com", "password")

	if lUser.ID != user.ID {
		t.Errorf("Ожидался User ID %d, получен %d", user.ID, lUser.ID)
	}
	if lToken.Value != token.Value {
		t.Errorf("Ожидался Token ID %s, получен %s", token.Value, lToken.Value)
	}
}
