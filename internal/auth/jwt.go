package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
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

// base64Encode encodes a byte slice into a base64 URL-encoded string without padding.
func base64Encode(input []byte) string {
	return strings.TrimRight(base64.URLEncoding.EncodeToString(input), "=")
}

// Decode base64 URL-encoded string and remove padding.
func base64Decode(input string) ([]byte, error) {
	paddedInput := input + strings.Repeat("=", (4-len(input)%4)%4)
	return base64.URLEncoding.DecodeString(paddedInput)
}

// createHMAC generates a base64 URL-encoded HMAC for a given message and secret key using SHA-256.
func createHMAC(message, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	return base64Encode(h.Sum(nil))
}

// GenerateJWT creates a JWT for a given user ID, username, and email.
// It returns the JWT as a string and an error if any occurs during the process.
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

// ValidateJWT validates the given JWT token
func ValidateJWT(token string) (*Claims, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, errors.New("invalid token format")
	}

	encodedHeader, encodedClaims, signature := parts[0], parts[1], parts[2]

	signingInput := fmt.Sprintf("%s.%s", encodedHeader, encodedClaims)
	expectedSignature := createHMAC(signingInput, jwtSecretKey)

	if signature != expectedSignature {
		return nil, errors.New("invalid token signature")
	}

	headerJSON, err := base64Decode(encodedHeader)
	if err != nil {
		return nil, fmt.Errorf("error decoding header: %v", err)
	}

	var header Header
	if err := json.Unmarshal(headerJSON, &header); err != nil {
		return nil, fmt.Errorf("error unmarshalling header: %v", err)
	}

	if header.Alg != "HS256" || header.Typ != "JWT" {
		return nil, errors.New("invalid token header")
	}

	claimsJSON, err := base64Decode(encodedClaims)
	if err != nil {
		return nil, fmt.Errorf("error decoding claims: %v", err)
	}

	var claims Claims
	if err := json.Unmarshal(claimsJSON, &claims); err != nil {
		return nil, fmt.Errorf("error unmarshalling claims: %v", err)
	}

	if claims.Exp < time.Now().Unix() {
		return nil, errors.New("token has expired")
	}

	return &claims, nil
}
