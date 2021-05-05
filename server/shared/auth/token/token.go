package token

import (
	"crypto/rsa"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type JWTTokenVerifier struct {
	PublicKey *rsa.PublicKey
}

func (v *JWTTokenVerifier) Verify(token string) (string, error) {
	claimsToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(*jwt.Token) (interface{}, error) {
		return v.PublicKey, nil
	})
	if err != nil {
		return "", fmt.Errorf("cannot parse token : %v ", err)
	}
	if !claimsToken.Valid {
		return "", fmt.Errorf("token no valid : %s ", token)
	}
	clm, ok := claimsToken.Claims.(*jwt.StandardClaims)

	if !ok {
		return "", fmt.Errorf("claims is not a standard claims %s ", token)
	}

	if err := clm.Valid(); err != nil {
		return "", fmt.Errorf("claims cannot be valid %s ", token)
	}

	return clm.Subject, nil
}
