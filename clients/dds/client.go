package dds

import (
	"github.com/morlay/aliyun-go/core"
)

func NewDdsClient(key string, secret string, regionId string) *DdsClient {
	return &DdsClient{
		Client: core.Client{
			Endpoint:        "https://dds.aliyuncs.com",
			Version:         "2015-12-01",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type DdsClient struct {
	core.Client
}
