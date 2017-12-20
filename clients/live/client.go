package live

import (
	"github.com/morlay/aliyun-go/core"
)

func NewLiveClient(key string, secret string, regionId string) *LiveClient {
	return &LiveClient{
		Client: core.Client{
			Endpoint:        "https://live.aliyuncs.com",
			Version:         "2016-11-01",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type LiveClient struct {
	core.Client
}
