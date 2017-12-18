package jaq

import "github.com/morlay/aliyun-go/core"

func NewJaqClient(key string, secret string, regionId string) *JaqClient {
	return &JaqClient{
		Client: core.Client{
			Endpoint:        "https://jaq.aliyuncs.com",
			Version:         "2016-11-23",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type JaqClient struct {
	core.Client
}
