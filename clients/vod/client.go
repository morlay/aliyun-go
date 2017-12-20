package vod

import (
	"github.com/morlay/aliyun-go/core"
)

func NewVodClient(key string, secret string, regionId string) *VodClient {
	return &VodClient{
		Client: core.Client{
			Endpoint:        "https://vod.aliyuncs.com",
			Version:         "2017-03-21",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type VodClient struct {
	core.Client
}
