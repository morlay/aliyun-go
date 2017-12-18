package qualitycheck

import "github.com/morlay/aliyun-go/core"

func NewQualitycheckClient(key string, secret string, regionId string) *QualitycheckClient {
	return &QualitycheckClient{
		Client: core.Client{
			Endpoint:        "https://qualitycheck.aliyuncs.com",
			Version:         "2016-08-01",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type QualitycheckClient struct {
	core.Client
}
