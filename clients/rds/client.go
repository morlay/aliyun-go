package rds

import (
	"github.com/morlay/aliyun-go/core"
)

func NewRdsClient(key string, secret string, regionId string) *RdsClient {
	return &RdsClient{
		Client: core.Client{
			Endpoint:        "https://rds.aliyuncs.com",
			Version:         "2014-08-15",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type RdsClient struct {
	core.Client
}
