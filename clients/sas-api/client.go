package sas_api

import (
	"github.com/morlay/aliyun-go/core"
)

func NewSasApiClient(key string, secret string, regionId string) *SasApiClient {
	return &SasApiClient{
		Client: core.Client{
			Endpoint:        "https://sas-api.aliyuncs.com",
			Version:         "2017-07-05",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type SasApiClient struct {
	core.Client
}
