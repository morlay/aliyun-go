package cloudauth

import "github.com/morlay/aliyun-go/core"

func NewCloudauthClient(key string, secret string, regionId string) *CloudauthClient {
	return &CloudauthClient{
		Client: core.Client{
			Endpoint:        "https://cloudauth.aliyuncs.com",
			Version:         "2017-11-17",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type CloudauthClient struct {
	core.Client
}
