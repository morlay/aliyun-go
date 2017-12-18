package crm

import "github.com/morlay/aliyun-go/core"

func NewCrmClient(key string, secret string, regionId string) *CrmClient {
	return &CrmClient{
		Client: core.Client{
			Endpoint:        "https://crm.aliyuncs.com",
			Version:         "2015-03-24",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type CrmClient struct {
	core.Client
}
