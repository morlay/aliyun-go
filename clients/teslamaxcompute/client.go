package teslamaxcompute

import (
	"github.com/morlay/aliyun-go/core"
)

func NewTeslamaxcomputeClient(key string, secret string, regionId string) *TeslamaxcomputeClient {
	return &TeslamaxcomputeClient{
		Client: core.Client{
			Endpoint:        "https://teslamaxcompute.aliyuncs.com",
			Version:         "2017-11-30",
			RegionID:        regionId,
			AccessKeyId:     key,
			AccessKeySecret: secret,
		},
	}
}

type TeslamaxcomputeClient struct {
	core.Client
}
