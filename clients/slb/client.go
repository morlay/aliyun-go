package slb

import (
	"github.com/morlay/aliyun-go/core"
)

func NewSlbClient(key string, secret string, regionId string) *SlbClient {
	return &SlbClient{
		Client: core.Client{
			Endpoint:        "https://slb.aliyuncs.com",
			Version:         "2014-05-15",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type SlbClient struct {
	core.Client
}
