package handlers

import (
	"testing"
)

func TestLoginHandler_Success(t *testing.T) {

	//TODO
	// add request validation test

	//// Генерируем случайную соль
	//salt := utils.GetSalt()
	//
	//// Хешируем пароль
	//hashedPassword := utils.HashPass(salt, "password")
	//user := &models.User{
	//	ID:       1,
	//	Email:    "test@example.com",
	//	Password: hex.EncodeToString(hashedPassword), // Сохраняем хеш (соль + хеш)
	//}
	//repo := &repository.MockUserRepository{User: user}
	//
	//token := utils.GenToken()
	//access := &models.Access{
	//	UserID: user.ID,
	//	Token:  token,
	//}
	//accessRepo := &repository.MockAccessRepository{Access: access}
	//
	//// Создаём запрос
	//reqBody := `{"email": "test@example.com", "password": "password"}`
	//req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(reqBody))
	//req.Header.Set("Content-Type", "application/json")
	//
	//rec := httptest.NewRecorder()
	//handler := LoginHandler(auth.NewAuthUseCase(repo, accessRepo))
	//handler.ServeHTTP(rec, req)
	//
	//// Проверяем статус код
	//if rec.Code != http.StatusOK {
	//	t.Errorf("Ожидался статус %d, получен %d", http.StatusOK, rec.Code)
	//}
	//
	//// Проверяем тело ответа
	//var resp LoginResponse
	//if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
	//	t.Fatalf("Ошибка декодирования ответа: %v", err)
	//}
	//if resp.ID != user.ID {
	//	t.Errorf("Ожидался ID %d, получен %d", user.ID, resp.ID)
	//}
}
