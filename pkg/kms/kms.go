package kms

import (
	"context"
	"errors"
	"strings"

	kmsGCP "cloud.google.com/go/kms/apiv1"
	kmsAWS "github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/octo-sts/app/pkg/gcpkms"
)

const (
	GCP = "gcp"
	AWS = "aws"
)

type kmsProvider struct {
	ctx       context.Context
	provider  string
	kmsKey    string
	gcpClient *kmsGCP.KeyManagementClient
	awsClient *kmsAWS.Client
}

type KMS interface {
	NewSigner() (ghinstallation.Signer, error)
}

// NewKMS provides a kmsProvider abstraction to allow for handling multiple providers and simplify passing around data
func NewKMS(ctx context.Context, provider, kmsKey string) (KMS, error) {
	kmsClient := &kmsProvider{
		ctx:      ctx,
		provider: strings.ToLower(provider),
		kmsKey:   kmsKey,
	}
	switch strings.ToLower(provider) {
	case GCP:
		gcpClient, err := kmsGCP.NewKeyManagementClient(ctx)
		if err != nil {
			return nil, err
		}
		kmsClient.gcpClient = gcpClient
		return kmsClient, nil
	case AWS:
		return nil, errors.New("not implemented")
	default:
		return nil, errors.New("unsupported kms provider")
	}
}

func (k *kmsProvider) NewSigner() (ghinstallation.Signer, error) {
	switch k.provider {
	case GCP:
		return gcpkms.New(k.ctx, k.gcpClient, k.kmsKey)
	case AWS:
		return nil, errors.New("not implemented")
	default:
		return nil, errors.New("unsupported kms provider")
	}
}
