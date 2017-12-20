package cloudphoto

import (
	"github.com/morlay/aliyun-go/core"
)

func NewCloudphotoClient(key string, secret string, regionId string) *CloudphotoClient {
	return &CloudphotoClient{
		Client: core.Client{
			Endpoint:        "https://cloudphoto.aliyuncs.com",
			Version:         "2017-07-11",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type CloudphotoClient struct {
	core.Client
}
