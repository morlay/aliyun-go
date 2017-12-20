package bss

import (
	"github.com/morlay/aliyun-go/core"
)

func NewBssClient(key string, secret string, regionId string) *BssClient {
	return &BssClient{
		Client: core.Client{
			Endpoint:        "https://bss.aliyuncs.com",
			Version:         "2014-07-14",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type BssClient struct {
	core.Client
}
