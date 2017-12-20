package cloudwf

import (
	"github.com/morlay/aliyun-go/core"
)

func NewCloudwfClient(key string, secret string, regionId string) *CloudwfClient {
	return &CloudwfClient{
		Client: core.Client{
			Endpoint:        "https://cloudwf.aliyuncs.com",
			Version:         "2017-03-28",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type CloudwfClient struct {
	core.Client
}
