package sts

import (
	"github.com/morlay/aliyun-go/core"
)

func NewStsClient(key string, secret string, regionId string) *StsClient {
	return &StsClient{
		Client: core.Client{
			Endpoint:        "https://sts.aliyuncs.com",
			Version:         "2015-04-01",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type StsClient struct {
	core.Client
}
