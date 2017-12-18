package hpc

import "github.com/morlay/aliyun-go/core"

func NewHpcClient(key string, secret string, regionId string) *HpcClient {
	return &HpcClient{
		Client: core.Client{
			Endpoint:        "https://hpc.aliyuncs.com",
			Version:         "2016-12-13",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type HpcClient struct {
	core.Client
}
