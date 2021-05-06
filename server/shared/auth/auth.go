package auth

import (
	"context"
	"coolcar/server/shared/auth/token"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"os"
	"strings"
)

const (
	accessTokenHeader = "authorization"
	bearerPrefix      = "Bearer "
)

type interceptor struct {
	verifyKey tokenVerify
}

type tokenVerify interface {
	Verify(token string) (string, error)
}

func Interceptor(pubKeyFile string) (grpc.UnaryServerInterceptor, error) {
	f, err := os.Open(pubKeyFile)
	if err != nil {
		return nil, fmt.Errorf("file not found %v", err)
	}
	key, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("file not found %v", err)
	}
	pem, _ := jwt.ParseRSAPublicKeyFromPEM(key)

	i := &interceptor{
		verifyKey: &token.JWTTokenVerifier{
			PublicKey: pem,
		},
	}

	return i.handleReq, nil
}

func (in *interceptor) handleReq(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	fromContext, err := tokenFromContext(ctx)

	aid, err := in.verifyKey.Verify(fromContext)

	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "fromContext error &v", err)
	}

	return handler(CtxWithAid(ctx, aid), req)
}

type accountIDKey struct{}

func CtxWithAid(c context.Context, aid string) context.Context {
	return context.WithValue(c, accountIDKey{}, aid)
}

func CidFromContext(c context.Context) (string, error) {
	v := c.Value(accountIDKey{})
	aid, ok := v.(string)
	if !ok {
		return "", status.Error(codes.Unauthenticated, "")
	}
	return aid, nil
}

func tokenFromContext(ctx context.Context) (string, error) {
	var t string
	if m, ok := metadata.FromIncomingContext(ctx); !ok {
		return "", status.Error(codes.Unauthenticated, "t error")
	} else {
		for _, val := range m[accessTokenHeader] {
			if strings.HasPrefix(val, bearerPrefix) {
				t = val[len(bearerPrefix):]
			}

		}
	}

	if t == "" {
		return "", status.Error(codes.Unauthenticated, "t error")
	}

	return t, nil
}
