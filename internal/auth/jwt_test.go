package auth

import (
	"encoding/base64"
	"encoding/json"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/joangavelan/contacts-app/config"
)

func TestGenerateJWT(t *testing.T) {
	// Set up environment variable for JWT_SECRET_KEY
	jwtSecretKey = "test_secret_key"
	os.Setenv("JWT_SECRET_KEY", jwtSecretKey)

	userId := int64(1)
	username := "testuser"
	email := "testuser@example.com"

	token, err := GenerateJWT(userId, username, email)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		t.Fatalf("expected token to have 3 parts, got %d", len(parts))
	}

	// Validate Header
	headerJSON, err := base64.URLEncoding.DecodeString(parts[0])
	if err != nil {
		t.Fatalf("error decoding header: %v", err)
	}
	var header Header
	if err := json.Unmarshal(headerJSON, &header); err != nil {
		t.Fatalf("error unmarshalling header: %v", err)
	}
	if header.Alg != "HS256" || header.Typ != "JWT" {
		t.Fatalf("unexpected header: %+v", header)
	}

	// Validate Claims
	claimsJSON, err := base64.URLEncoding.DecodeString(parts[1])
	if err != nil {
		t.Fatalf("error decoding claims: %v", err)
	}
	var claims Claims
	if err := json.Unmarshal(claimsJSON, &claims); err != nil {
		t.Fatalf("error unmarshalling claims: %v", err)
	}
	if claims.Sub != userId || claims.Email != email || claims.Username != username {
		t.Fatalf("unexpected claims: %+v", claims)
	}
	now := time.Now().Unix()
	if claims.Iat < now-1 || claims.Iat > now+1 {
		t.Fatalf("unexpected iat claim: %d", claims.Iat)
	}
	if claims.Exp != claims.Iat+int64(config.JWTExpiration.Seconds()) {
		t.Fatalf("unexpected exp claim: %d", claims.Exp)
	}

	// Validate Signature
	signingInput := parts[0] + "." + parts[1]
	expectedSignature := createHMAC(signingInput, jwtSecretKey)
	if parts[2] != expectedSignature {
		t.Fatalf("unexpected signature: %s, expected: %s", parts[2], expectedSignature)
	}
}

func TestValidateJWT(t *testing.T) {
	// Setup
	jwtSecretKey = "testsecretkey"

	// Generate a valid token
	userID := int64(1)
	username := "testuser"
	email := "testuser@example.com"
	token, err := GenerateJWT(userID, username, email)
	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}

	// Test valid token
	claims, err := ValidateJWT(token)
	if err != nil {
		t.Errorf("expected valid token, got error: %v", err)
	}
	if claims.Sub != userID || claims.Username != username || claims.Email != email {
		t.Errorf("claims do not match expected values: got %+v", claims)
	}

	// Test invalid token signature
	invalidSignatureToken := token[:len(token)-1] + "X"
	_, err = ValidateJWT(invalidSignatureToken)
	if err == nil || err.Error() != "invalid token signature" {
		t.Errorf("expected invalid token signature error, got: %v", err)
	}

	// Test expired token
	expiredClaims := Claims{
		Sub:      userID,
		Iat:      time.Now().Add(-2 * time.Hour).Unix(),
		Exp:      time.Now().Add(-1 * time.Hour).Unix(),
		Email:    email,
		Username: username,
	}
	expiredClaimsJSON, _ := json.Marshal(expiredClaims)
	expiredToken := strings.Split(token, ".")[0] + "." + base64Encode(expiredClaimsJSON) + "." + createHMAC(strings.Split(token, ".")[0]+"."+base64Encode(expiredClaimsJSON), jwtSecretKey)
	_, err = ValidateJWT(expiredToken)
	if err == nil || err.Error() != "token has expired" {
		t.Errorf("expected token has expired error, got: %v", err)
	}
}
