package mts

import "github.com/morlay/aliyun-go/core"

func NewMtsClient(key string, secret string, regionId string) *MtsClient {
	return &MtsClient{
		Client: core.Client{
			Endpoint:        "https://mts.aliyuncs.com",
			Version:         "2014-06-18",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type MtsClient struct {
	core.Client
}
