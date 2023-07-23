package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func ComparePasswords(hashedpassword string, plainpassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(plainpassword))
	if err != nil {
		fmt.Printf("bcrypt error: %v", err)
		return false
	}
	return true
}

func GenerateToken(userid int) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Ejemplo: expira en 24 horas

	// Crea los claims (datos) del token
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Id:        fmt.Sprintf("%d", userid),
	}

	// creacion de token con claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// clave secreta
	jwtkey := []byte("tengohambre")

	// token jwt firmado con la clave
	tokenstring, err := token.SignedString(jwtkey)
	if err != nil {
		return "", err
	}

	return tokenstring, nil
}
