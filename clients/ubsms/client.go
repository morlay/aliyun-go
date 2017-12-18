package ubsms

import "github.com/morlay/aliyun-go/core"

func NewUbsmsClient(key string, secret string, regionId string) *UbsmsClient {
	return &UbsmsClient{
		Client: core.Client{
			Endpoint:        "https://ubsms.aliyuncs.com",
			Version:         "2015-06-23",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type UbsmsClient struct {
	core.Client
}
