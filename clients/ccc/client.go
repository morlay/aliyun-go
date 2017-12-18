package ccc

import "github.com/morlay/aliyun-go/core"

func NewCccClient(key string, secret string, regionId string) *CccClient {
	return &CccClient{
		Client: core.Client{
			Endpoint:        "https://ccc.aliyuncs.com",
			Version:         "2017-07-05",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type CccClient struct {
	core.Client
}
