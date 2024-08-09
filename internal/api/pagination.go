package api

import (
	"encoding/base64"
	"fmt"

	"google.golang.org/protobuf/proto"

	pbv1 "github.com/invzhi/outward/proto/outward/v1"
)

func NewPageToken(pageToken *pbv1.PageToken) (string, error) {
	b, err := proto.Marshal(pageToken)
	if err != nil {
		return "", fmt.Errorf("cannot marshal page token: %w", err)
	}

	return base64.StdEncoding.EncodeToString(b), nil
}

func ParsePageToken(pageToken string) (*pbv1.PageToken, error) {
	b, err := base64.StdEncoding.DecodeString(pageToken)
	if err != nil {
		return nil, fmt.Errorf("cannot decode page token: %w", err)
	}

	var token pbv1.PageToken
	err = proto.Unmarshal(b, &token)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal page token: %w", err)
	}
	return &token, nil
}
