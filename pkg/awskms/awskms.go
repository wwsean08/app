package awskms

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/golang-jwt/jwt/v4"
)

type signingMethodAWS struct {
	client *kms.Client
}

func (s *signingMethodAWS) Verify(signingString, signature string, key interface{}) error {
	return errors.New("not implemented")
}

func (s *signingMethodAWS) Sign(signingString string, key interface{}) (string, error) {
	return "", errors.New("not implemented")
}

func (s *signingMethodAWS) Alg() string {
	return ""
}

type awsSigner struct {
	client *kms.Client
	key    string
}

func New(_ context.Context, client *kms.Client, key string) (ghinstallation.Signer, error) {
	return nil, errors.New("not implemented")
}

func (s *awsSigner) Sign(claims jwt.Claims) (string, error) {
	return "", errors.New("not implemented")
}
