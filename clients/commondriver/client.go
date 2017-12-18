package commondriver

import "github.com/morlay/aliyun-go/core"

func NewCommondriverClient(key string, secret string, regionId string) *CommondriverClient {
	return &CommondriverClient{
		Client: core.Client{
			Endpoint:        "https://commondriver.aliyuncs.com",
			Version:         "2015-12-29",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type CommondriverClient struct {
	core.Client
}
