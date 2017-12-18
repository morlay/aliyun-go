package domain

import "github.com/morlay/aliyun-go/core"

func NewDomainClient(key string, secret string, regionId string) *DomainClient {
	return &DomainClient{
		Client: core.Client{
			Endpoint:        "https://domain.aliyuncs.com",
			Version:         "2016-05-11",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type DomainClient struct {
	core.Client
}
