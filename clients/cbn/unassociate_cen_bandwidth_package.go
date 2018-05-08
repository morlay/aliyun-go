package cbn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type UnassociateCenBandwidthPackageRequest struct {
	requests.RpcRequest
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	CenId                 string `position:"Query" name:"CenId"`
	CenBandwidthPackageId string `position:"Query" name:"CenBandwidthPackageId"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
}

func (req *UnassociateCenBandwidthPackageRequest) Invoke(client *sdk.Client) (resp *UnassociateCenBandwidthPackageResponse, err error) {
	req.InitWithApiInfo("Cbn", "2017-09-12", "UnassociateCenBandwidthPackage", "cbn", "")
	resp = &UnassociateCenBandwidthPackageResponse{}
	err = client.DoAction(req, resp)
	return
}

type UnassociateCenBandwidthPackageResponse struct {
	responses.BaseResponse
	RequestId common.String
}
