package cloudapi

import (
	"github.com/morlay/aliyun-go/core"
)

func NewCloudapiClient(key string, secret string, regionId string) *CloudapiClient {
	return &CloudapiClient{
		Client: core.Client{
			Endpoint:        "https://cloudapi.aliyuncs.com",
			Version:         "2016-07-14",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type CloudapiClient struct {
	core.Client
}
