package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/joangavelan/contacts-app/config"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Claims struct {
	Sub      int64  `json:"sub"`
	Exp      int64  `json:"exp"`
	Iat      int64  `json:"iat"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

var jwtSecretKey = os.Getenv("JWT_SECRET_KEY")

func base64Encode(input []byte) string {
	return strings.TrimRight(base64.URLEncoding.EncodeToString(input), "=")
}

func createHMAC(message, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	return base64Encode(h.Sum(nil))
}

func GenerateJWT(userId int64, username, email string) (string, error) {
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}

	headerJSON, err := json.Marshal(header)
	if err != nil {
		return "", fmt.Errorf("error encoding header: %v", err)
	}

	now := time.Now().Unix()
	claims := Claims{
		Sub:      userId,
		Iat:      now,
		Exp:      now + int64(config.JWTExpiration.Seconds()),
		Email:    email,
		Username: username,
	}

	claimsJSON, err := json.Marshal(claims)
	if err != nil {
		return "", fmt.Errorf("error encoding claims %v", err)
	}

	encodedHeader := base64Encode(headerJSON)
	encodedClaims := base64Encode(claimsJSON)

	signingInput := fmt.Sprintf("%s.%s", encodedHeader, encodedClaims)
	signature := createHMAC(signingInput, jwtSecretKey)

	token := fmt.Sprintf("%s.%s.%s", encodedHeader, encodedClaims, signature)

	return token, nil
}
