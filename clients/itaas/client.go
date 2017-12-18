package itaas

import "github.com/morlay/aliyun-go/core"

func NewItaasClient(key string, secret string, regionId string) *ItaasClient {
	return &ItaasClient{
		Client: core.Client{
			Endpoint:        "https://itaas.aliyuncs.com",
			Version:         "2017-05-05",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type ItaasClient struct {
	core.Client
}
