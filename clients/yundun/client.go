package yundun

import (
	"github.com/morlay/aliyun-go/core"
)

func NewYundunClient(key string, secret string, regionId string) *YundunClient {
	return &YundunClient{
		Client: core.Client{
			Endpoint:        "https://yundun.aliyuncs.com",
			Version:         "2015-04-16",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type YundunClient struct {
	core.Client
}
