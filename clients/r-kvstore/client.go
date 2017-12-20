package r_kvstore

import (
	"github.com/morlay/aliyun-go/core"
)

func NewRKvstoreClient(key string, secret string, regionId string) *RKvstoreClient {
	return &RKvstoreClient{
		Client: core.Client{
			Endpoint:        "https://r-kvstore.aliyuncs.com",
			Version:         "2015-01-01",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type RKvstoreClient struct {
	core.Client
}
