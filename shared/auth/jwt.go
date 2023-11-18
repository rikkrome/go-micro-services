package auth

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthTokens struct {
	AccessToken        string
	AccessTokenExpiry  int64
	RefreshToken       string
	RefreshTokenExpiry int64
}

func NewToken(id string) (tokens AuthTokens, err error) {
	// Create an array to hold the key
	base64Key := os.Getenv("SECRET_AUTH_KEY")
	// Decode the Base64 encoded key
	key, err := base64.StdEncoding.DecodeString(base64Key)
	if err != nil {
		log.Fatal("Error decoding Base64 key: ", err)
		return AuthTokens{}, err
	}
	// Token
	tokenExp := time.Now().Add(time.Hour * 72).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": tokenExp,
	})
	tokenStr, err := token.SignedString(key)
	if err != nil {
		fmt.Println("Error SignedString:", err)
		return
	}
	// Refresh Token
	refreshTokenExp := time.Now().Add(time.Hour * 24 * 30).Unix()
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": refreshTokenExp,
	})
	refreshTokenStr, err := refreshToken.SignedString(key)
	if err != nil {
		fmt.Println("Error SignedString:", err)
		return
	}
	tokens = AuthTokens{
		AccessToken:        tokenStr,
		AccessTokenExpiry:  tokenExp,
		RefreshToken:       refreshTokenStr,
		RefreshTokenExpiry: refreshTokenExp,
	}
	return tokens, nil
}
