package ons

import (
	"github.com/morlay/aliyun-go/core"
)

func NewOnsClient(key string, secret string, regionId string) *OnsClient {
	return &OnsClient{
		Client: core.Client{
			Endpoint:        "https://ons.aliyuncs.com",
			Version:         "2017-09-18",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type OnsClient struct {
	core.Client
}
