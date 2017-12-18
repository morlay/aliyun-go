package emr

import "github.com/morlay/aliyun-go/core"

func NewEmrClient(key string, secret string, regionId string) *EmrClient {
	return &EmrClient{
		Client: core.Client{
			Endpoint:        "https://emr.aliyuncs.com",
			Version:         "2016-04-08",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type EmrClient struct {
	core.Client
}
