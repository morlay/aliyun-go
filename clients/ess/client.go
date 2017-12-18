package ess

import "github.com/morlay/aliyun-go/core"

func NewEssClient(key string, secret string, regionId string) *EssClient {
	return &EssClient{
		Client: core.Client{
			Endpoint:        "https://ess.aliyuncs.com",
			Version:         "2014-08-28",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type EssClient struct {
	core.Client
}
