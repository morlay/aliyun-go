package oms

import "github.com/morlay/aliyun-go/core"

func NewOmsClient(key string, secret string, regionId string) *OmsClient {
	return &OmsClient{
		Client: core.Client{
			Endpoint:        "https://oms.aliyuncs.com",
			Version:         "2015-02-12",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type OmsClient struct {
	core.Client
}
