package nas

import (
	"github.com/morlay/aliyun-go/core"
)

func NewNasClient(key string, secret string, regionId string) *NasClient {
	return &NasClient{
		Client: core.Client{
			Endpoint:        "https://nas.aliyuncs.com",
			Version:         "2017-06-26",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type NasClient struct {
	core.Client
}
