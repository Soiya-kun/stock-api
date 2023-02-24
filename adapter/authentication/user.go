package authentication

import (
	"fmt"
	"time"

	"gitlab.com/soy-app/stock-api/config"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

var userTokenJwt = &hs256jwt{
	sigKey: []byte(config.SigKey()),
	createClaims: func() jwt.Claims {
		return &userClaims{}
	},
}

type userClaims struct {
	ID       string `json:"jti"`
	Subject  string `json:"sub"`
	IssuedAt int    `json:"iat"`
}

func (c *userClaims) Valid() error {
	now := time.Now()
	if c.IssuedAt > int(now.Unix()) {
		return fmt.Errorf("issued on future time: %d (now:%d)", c.IssuedAt, now.Unix())
	}

	_, err := uuid.Parse(c.ID)
	if err != nil {
		return fmt.Errorf("invalid id=%s: %w", c.ID, err)
	}

	return nil
}

func IssueUserToken(userID string) (string, error) {
	id := uuid.New()
	now := time.Now()

	claims := &userClaims{
		ID:       id.String(),
		Subject:  userID,
		IssuedAt: int(now.Unix()),
	}

	return userTokenJwt.issueToken(claims)
}

func VerifyUserToken(token string) (string, error) {
	claims, err := userTokenJwt.verifyToken(token)
	if err != nil {
		return "", err
	}

	return claims.(*userClaims).Subject, nil
}
