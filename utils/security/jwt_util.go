package security

import (
	"fmt"
	"polen/config"
	"polen/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwtToken(user model.UserCredential) (string, error) {
	now := time.Now()
	end := now.Add(time.Duration(60) * time.Minute)

	claims := &AppClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "polen",
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(end),
		},
		Username: user.Username,
		// Role:     "",
		// Services: []string{},
	}

	tokenJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Ini kita pindah ke config ENV
	tokenString, err := tokenJwt.SignedString([]byte("secret"))
	if err != nil {
		return "", fmt.Errorf("failed to create jwt token: %v", err.Error())
	}
	return tokenString, nil
}

func VerifyJwtToken(tokenString string) (jwt.MapClaims, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		method, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok || method != cfg.JwtSigningMethod {
			return nil, fmt.Errorf("invalid token signin method")
		}
		return cfg.JwtSignatureKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error: %v", err.Error())
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid || claims["iss"] != cfg.ApplicationName {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}
