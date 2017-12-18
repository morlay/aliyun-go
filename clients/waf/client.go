package waf

import "github.com/morlay/aliyun-go/core"

func NewWafClient(key string, secret string, regionId string) *WafClient {
	return &WafClient{
		Client: core.Client{
			Endpoint:        "https://waf.aliyuncs.com",
			Version:         "2016-11-11",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type WafClient struct {
	core.Client
}
