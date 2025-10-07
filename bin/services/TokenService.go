package services

import (
	"dochub/bin/models"
	"math/rand"
	"strings"
	"time"
)

func GenerateToken(user models.User) *models.UserToken {
	length := 10 // Desired length of the random string
	randomString := GenerateRandomString(length)

	now := time.Now()
	twoDaysLater := now.Add(2 * 24 * time.Hour) // add 2 days

	token := &models.UserToken{
		Email:     user.Email,
		Token:     randomString,
		ValidTill: &twoDaysLater,
	}

	return token
}

func GenerateRandomString(m int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano())) // Seed for older Go versions

	sb := strings.Builder{}
	sb.Grow(m) // Pre-allocate memory for efficiency
	for i := 0; i < m; i++ {
		sb.WriteByte(charset[seededRand.Intn(len(charset))])
	}
	return sb.String()
}
