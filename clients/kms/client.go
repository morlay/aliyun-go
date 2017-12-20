package kms

import (
	"github.com/morlay/aliyun-go/core"
)

func NewKmsClient(key string, secret string, regionId string) *KmsClient {
	return &KmsClient{
		Client: core.Client{
			Endpoint:        "https://kms.aliyuncs.com",
			Version:         "2016-01-20",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type KmsClient struct {
	core.Client
}
