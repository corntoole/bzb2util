package backblaze

import (
	"github.com/pkg/errors"

	"github.com/GiantRavens/backblazeS3/backblazeS3"
)

var (
	b2               backblazeS3.B2
	ErrUnimplemented = errors.New("method is not implemnted")
)

type B2Client struct {
	backblazeS3.B2Client
}

var _ backblazeS3.B2 = &B2Client{}

func NewB2Client(endpoint, region, keyID, applicationKey, token, bucketName string) (backblazeS3.B2, error) {
	return backblazeS3.NewB2Client(endpoint, region, keyID, applicationKey, token, bucketName)
}

func (b *B2Client) Delete(_ string) error {
	return ErrUnimplemented
}
