package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AddBandwidthPackageIpsRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	BandwidthPackageId   string `position:"Query" name:"BandwidthPackageId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	IpCount              string `position:"Query" name:"IpCount"`
}

func (req *AddBandwidthPackageIpsRequest) Invoke(client *sdk.Client) (resp *AddBandwidthPackageIpsResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "AddBandwidthPackageIps", "ecs", "")
	resp = &AddBandwidthPackageIpsResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddBandwidthPackageIpsResponse struct {
	responses.BaseResponse
	RequestId common.String
}
