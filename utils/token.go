package utils

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"strconv"
	"time"
)

func GenerateToken(userID uint, roles []string) (string, error) {
	expirationHoursStr := os.Getenv("TOKEN_HOUR_LIFESPAN")
	expirationHours, err := strconv.Atoi(expirationHoursStr)
	if err != nil {
		// Default to 24 hours if TOKEN_HOUR_LIFESPAN is not set or invalid
		expirationHours = 24
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"roles":   roles,
		"exp":     time.Now().Add(time.Hour * time.Duration(expirationHours)).Unix(),
	})

	apiSecret := os.Getenv("API_SECRET")
	// Sign and get the complete encoded token as a string using the secret key
	return token.SignedString([]byte(apiSecret))
}
