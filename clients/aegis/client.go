package aegis

import "github.com/morlay/aliyun-go/core"

func NewAegisClient(key string, secret string, regionId string) *AegisClient {
	return &AegisClient{
		Client: core.Client{
			Endpoint:        "https://aegis.aliyuncs.com",
			Version:         "2016-11-11",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type AegisClient struct {
	core.Client
}
