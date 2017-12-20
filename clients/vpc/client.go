package vpc

import (
	"github.com/morlay/aliyun-go/core"
)

func NewVpcClient(key string, secret string, regionId string) *VpcClient {
	return &VpcClient{
		Client: core.Client{
			Endpoint:        "https://vpc.aliyuncs.com",
			Version:         "2016-04-28",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type VpcClient struct {
	core.Client
}
