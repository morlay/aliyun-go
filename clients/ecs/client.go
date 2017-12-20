package ecs

import (
	"github.com/morlay/aliyun-go/core"
)

func NewEcsClient(key string, secret string, regionId string) *EcsClient {
	return &EcsClient{
		Client: core.Client{
			Endpoint:        "https://ecs.aliyuncs.com",
			Version:         "2014-05-26",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type EcsClient struct {
	core.Client
}
