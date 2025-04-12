// Copyright 2024 Chainguard, Inc.
// SPDX-License-Identifier: Apache-2.0

package awskms

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/golang-jwt/jwt/v4"
)

type signingMethodAWS struct {
	client *kms.Client
}

func (s *signingMethodAWS) Verify(signingString, signature string, key interface{}) error {
	return errors.New("not implemented")
}

func (s *signingMethodAWS) Sign(signingString string, ikey interface{}) (string, error) {
	println(signingString)
	ctx := context.Background()

	key, ok := ikey.(string)
	if !ok {
		return "", fmt.Errorf("invalid key reference type: %T", ikey)
	}
	resp, err := s.client.Sign(ctx, &kms.SignInput{
		KeyId:            aws.String(key),
		Message:          []byte(signingString),
		MessageType:      types.MessageTypeRaw,
		SigningAlgorithm: types.SigningAlgorithmSpecRsassaPkcs1V15Sha256,
	})
	if err != nil {
		return "", err
	}

	println(base64.RawURLEncoding.EncodeToString(resp.Signature))
	return base64.RawURLEncoding.EncodeToString(resp.Signature), errors.New("not implemented")
}

func (s *signingMethodAWS) Alg() string {
	return "RS256"
}

type awsSigner struct {
	client *kms.Client
	key    string
}

func New(_ context.Context, client *kms.Client, key string) (ghinstallation.Signer, error) {
	return &awsSigner{
		client: client,
		key:    key,
	}, nil
}

func (s *awsSigner) Sign(claims jwt.Claims) (string, error) {
	method := &signingMethodAWS{
		client: s.client,
	}
	return jwt.NewWithClaims(method, claims).SignedString(s.key)
}
