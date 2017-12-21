package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyBandwidthPackageSpecRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	BandwidthPackageId   string `position:"Query" name:"BandwidthPackageId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth            string `position:"Query" name:"Bandwidth"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifyBandwidthPackageSpecRequest) Invoke(client *sdk.Client) (resp *ModifyBandwidthPackageSpecResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifyBandwidthPackageSpec", "ecs", "")
	resp = &ModifyBandwidthPackageSpecResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyBandwidthPackageSpecResponse struct {
	responses.BaseResponse
	RequestId string
}
