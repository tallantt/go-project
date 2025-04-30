package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"time"
)

// Ваш секретный ключ для подписи токенов
var secret = []byte("your-secret-key")

// Функция для генерации JWT с ролью
func GenerateJWT(userID uint, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role, // Добавляем роль в claims
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

// Функция для валидации JWT и извлечения данных
func ValidateJWT(tokenStr string) (*jwt.Token, jwt.MapClaims, error) {
	// Разбираем токен
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Проверка метода подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secret, nil
	})

	// Если произошла ошибка при разборе токена или токен недействителен
	if err != nil || !token.Valid {
		return nil, nil, err
	}

	// Извлекаем claims из токена
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, nil, errors.New("invalid claims")
	}

	// Возвращаем токен, claims и ошибку (если она есть)
	return token, claims, nil
}
