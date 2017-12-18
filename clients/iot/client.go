package iot

import "github.com/morlay/aliyun-go/core"

func NewIotClient(key string, secret string, regionId string) *IotClient {
	return &IotClient{
		Client: core.Client{
			Endpoint:        "https://iot.aliyuncs.com",
			Version:         "2017-04-20",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type IotClient struct {
	core.Client
}
