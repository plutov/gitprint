package git

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/plutov/gitprint/api/pkg/log"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

type Auth struct {
	conf   *oauth2.Config
	states *TTLMap
}

type User struct {
	Email       string `json:"email"`
	Username    string `json:"username"`
	AccessToken string `json:"access_token"`
}

func NewAuth() *Auth {
	return &Auth{
		conf:   getOAuthConfig(),
		states: NewTTLMap(1000, 3600),
	}
}

func getOAuthConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		Scopes:       []string{"user:email", "repo"},
		Endpoint:     github.Endpoint,
	}
}

func getRandomState() string {
	timestamp := time.Now().UnixNano()

	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		log.WithError(err).Error("unable to generate random state")
		return fmt.Sprintf("%x", sha256.Sum256([]byte(strconv.Itoa(int(timestamp)))))
	}

	salt := base64.URLEncoding.EncodeToString(b)

	return fmt.Sprintf("%x", sha256.Sum256([]byte(salt+strconv.Itoa(int(timestamp)))))
}

func (a *Auth) GetRedirectURL() string {
	state := getRandomState()
	a.states.Put(state)

	return a.conf.AuthCodeURL(state)
}

func (a *Auth) GetAccessToken(code string, state string) (string, error) {
	logCtx := log.With("code", code, "state", state)
	logCtx.Info("getting access token")

	if ok := a.states.Ok(state); !ok {
		logCtx.Warn("invalid state")
		return "", fmt.Errorf("invalid state")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	token, err := a.conf.Exchange(ctx, code)
	if err != nil {
		logCtx.WithError(err).Error("exchange failed")
		return "", fmt.Errorf("exchange failed")
	}

	a.states.Delete(state)

	logCtx.Info("access token received")

	return token.AccessToken, nil
}
