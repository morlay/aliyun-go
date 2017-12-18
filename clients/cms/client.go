package cms

import "github.com/morlay/aliyun-go/core"

func NewCmsClient(key string, secret string, regionId string) *CmsClient {
	return &CmsClient{
		Client: core.Client{
			Endpoint:        "https://cms.aliyuncs.com",
			Version:         "2017-03-01",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type CmsClient struct {
	core.Client
}
