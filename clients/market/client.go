package market

import "github.com/morlay/aliyun-go/core"

func NewMarketClient(key string, secret string, regionId string) *MarketClient {
	return &MarketClient{
		Client: core.Client{
			Endpoint:        "https://market.aliyuncs.com",
			Version:         "2015-11-01",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type MarketClient struct {
	core.Client
}
