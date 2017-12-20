package dm

import (
	"github.com/morlay/aliyun-go/core"
)

func NewDmClient(key string, secret string, regionId string) *DmClient {
	return &DmClient{
		Client: core.Client{
			Endpoint:        "https://dm.aliyuncs.com",
			Version:         "2015-11-23",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type DmClient struct {
	core.Client
}
