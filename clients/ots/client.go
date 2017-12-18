package ots

import "github.com/morlay/aliyun-go/core"

func NewOtsClient(key string, secret string, regionId string) *OtsClient {
	return &OtsClient{
		Client: core.Client{
			Endpoint:        "https://ots.aliyuncs.com",
			Version:         "2013-09-12",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type OtsClient struct {
	core.Client
}
