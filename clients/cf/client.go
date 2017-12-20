package cf

import (
	"github.com/morlay/aliyun-go/core"
)

func NewCfClient(key string, secret string, regionId string) *CfClient {
	return &CfClient{
		Client: core.Client{
			Endpoint:        "https://cf.aliyuncs.com",
			Version:         "2015-12-08",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type CfClient struct {
	core.Client
}
