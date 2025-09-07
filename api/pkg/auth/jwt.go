package auth

import (
	"crypto/sha512"
	"errors"
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/plutov/gitprint/api/pkg/git"
)

var tempJwtSecret []byte

type SessionClaims struct {
	User *git.User `json:"user"`
	jwt.StandardClaims
}

func init() {
	var sha512 = sha512.New()
	sha512.Write([]byte(fmt.Sprintf("gitprintme%d", time.Now().UnixNano())))

	tempJwtSecret = sha512.Sum(nil)
}

func FillJWT(user *git.User) (string, error) {
	claims := &SessionClaims{
		User: user,
	}
	claims.ExpiresAt = time.Now().Add(time.Hour * 24 * 30).UTC().Unix()

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response
	t, err := token.SignedString(tempJwtSecret)
	if err != nil {
		return "", err
	}

	return t, nil
}

func ReadJWTClaims(jwtToken string) (*SessionClaims, error) {
	token, err := jwt.ParseWithClaims(jwtToken, &SessionClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return tempJwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*SessionClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid claims")
}
