package drds

import "github.com/morlay/aliyun-go/core"

func NewDrdsClient(key string, secret string, regionId string) *DrdsClient {
	return &DrdsClient{
		Client: core.Client{
			Endpoint:        "https://drds.aliyuncs.com",
			Version:         "2017-10-16",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type DrdsClient struct {
	core.Client
}
