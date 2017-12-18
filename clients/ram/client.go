package ram

import "github.com/morlay/aliyun-go/core"

func NewRamClient(key string, secret string, regionId string) *RamClient {
	return &RamClient{
		Client: core.Client{
			Endpoint:        "https://ram.aliyuncs.com",
			Version:         "2015-05-01",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type RamClient struct {
	core.Client
}
