package push

import (
	"github.com/morlay/aliyun-go/core"
)

func NewPushClient(key string, secret string, regionId string) *PushClient {
	return &PushClient{
		Client: core.Client{
			Endpoint:        "https://push.aliyuncs.com",
			Version:         "2016-08-01",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type PushClient struct {
	core.Client
}
