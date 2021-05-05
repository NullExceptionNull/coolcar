package auth

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWTTokenGen struct {
	Issuer     string
	nowFunc    func() time.Time
	privateKey *rsa.PrivateKey
}

func NewJWTTokenGen(issuer string, privateKey *rsa.PrivateKey) *JWTTokenGen {
	return &JWTTokenGen{Issuer: issuer,
		nowFunc:    time.Now,
		privateKey: privateKey}
}

func (t *JWTTokenGen) GenerateToken(accountID string, expireTime time.Duration) (string, error) {
	nowSec := t.nowFunc().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.StandardClaims{
		Issuer:    t.Issuer,
		IssuedAt:  nowSec,
		ExpiresAt: nowSec + int64(expireTime),
		Subject:   accountID,
	})
	return token.SignedString(t.privateKey)
}
