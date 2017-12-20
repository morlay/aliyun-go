package ocs

import (
	"github.com/morlay/aliyun-go/core"
)

func NewOcsClient(key string, secret string, regionId string) *OcsClient {
	return &OcsClient{
		Client: core.Client{
			Endpoint:        "https://ocs.aliyuncs.com",
			Version:         "2015-03-01",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type OcsClient struct {
	core.Client
}
