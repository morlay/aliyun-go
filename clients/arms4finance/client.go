package arms4finance

import (
	"github.com/morlay/aliyun-go/core"
)

func NewArms4financeClient(key string, secret string, regionId string) *Arms4financeClient {
	return &Arms4financeClient{
		Client: core.Client{
			Endpoint:        "https://arms4finance.aliyuncs.com",
			Version:         "2017-11-30",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type Arms4financeClient struct {
	core.Client
}
