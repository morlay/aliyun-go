package httpdns

import "github.com/morlay/aliyun-go/core"

func NewHttpdnsClient(key string, secret string, regionId string) *HttpdnsClient {
	return &HttpdnsClient{
		Client: core.Client{
			Endpoint:        "https://httpdns.aliyuncs.com",
			Version:         "2016-02-01",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type HttpdnsClient struct {
	core.Client
}
