package ossadmin

import "github.com/morlay/aliyun-go/core"

func NewOssadminClient(key string, secret string, regionId string) *OssadminClient {
	return &OssadminClient{
		Client: core.Client{
			Endpoint:        "https://ossadmin.aliyuncs.com",
			Version:         "2015-03-02",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type OssadminClient struct {
	core.Client
}
