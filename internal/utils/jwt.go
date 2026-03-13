	package utils

	import (
		"os"
		"time"
		"github.com/golang-jwt/jwt/v4"
	)

	var jwtKey = []byte(os.Getenv("JWT_SECRET"))

	type Claims struct {
		ID    int    `json:"id"`
		Email string `json:"email"`
		jwt.RegisteredClaims
	}

	func GenerateToken(ID int, email string) (string, error) {
		expirationTime := time.Now().Add(24 * time.Hour)
		claims := &Claims{
			ID  : ID,
			Email: email,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expirationTime),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		return token.SignedString(jwtKey)
	}

	func ValidateToken(tokenStr string) (*Claims, error) {
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			return nil, err
		}
		return claims, nil
	}
