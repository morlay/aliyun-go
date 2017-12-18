package alidns

import "github.com/morlay/aliyun-go/core"

func NewAlidnsClient(key string, secret string, regionId string) *AlidnsClient {
	return &AlidnsClient{
		Client: core.Client{
			Endpoint:        "https://alidns.aliyuncs.com",
			Version:         "2015-01-09",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type AlidnsClient struct {
	core.Client
}
