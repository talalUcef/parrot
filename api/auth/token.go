package auth

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

type AuthProvider struct {
	Name       string
	SigningKey []byte
}

func (ap *AuthProvider) CreateToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(ap.SigningKey)
}

func (ap *AuthProvider) ParseAndVerifyToken(tokenString string) (jwt.MapClaims, error) {
	return parseAndVerify(tokenString, ap.SigningKey)
}

func parseAndVerify(tokenString string, signingKey []byte) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return signingKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims.Valid() != nil {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}
