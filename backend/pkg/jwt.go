package pkg

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTGen struct {
	secret string
	expiry time.Duration
}

func NewJWTGen(secret, expiryMinutes string) (*JWTGen, error) {
	expiry, err := strconv.Atoi(expiryMinutes)
	if err != nil {
		return nil, err
	}

	return &JWTGen{
		secret: secret,
		expiry: time.Duration(expiry) * time.Minute,
	}, nil
}

type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func (j *JWTGen) Generate(userID uint, email string) (string, error) {
	claims := Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.expiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secret))
}

func (j *JWTGen) Validate(tokenStr string) (map[string]any, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(j.secret), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return map[string]any{
		"user_id": claims.UserID,
		"email":   claims.Email,
	}, nil
}
