package cbn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteCenBandwidthPackageRequest struct {
	requests.RpcRequest
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	CenBandwidthPackageId string `position:"Query" name:"CenBandwidthPackageId"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteCenBandwidthPackageRequest) Invoke(client *sdk.Client) (resp *DeleteCenBandwidthPackageResponse, err error) {
	req.InitWithApiInfo("Cbn", "2017-09-12", "DeleteCenBandwidthPackage", "cbn", "")
	resp = &DeleteCenBandwidthPackageResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteCenBandwidthPackageResponse struct {
	responses.BaseResponse
	RequestId common.String
}
