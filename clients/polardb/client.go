package polardb

import "github.com/morlay/aliyun-go/core"

func NewPolardbClient(key string, secret string, regionId string) *PolardbClient {
	return &PolardbClient{
		Client: core.Client{
			Endpoint:        "https://polardb.aliyuncs.com",
			Version:         "2017-08-01",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type PolardbClient struct {
	core.Client
}
