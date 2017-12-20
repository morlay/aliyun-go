package cdn

import (
	"github.com/morlay/aliyun-go/core"
)

func NewCdnClient(key string, secret string, regionId string) *CdnClient {
	return &CdnClient{
		Client: core.Client{
			Endpoint:        "https://cdn.aliyuncs.com",
			Version:         "2014-11-11",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type CdnClient struct {
	core.Client
}
