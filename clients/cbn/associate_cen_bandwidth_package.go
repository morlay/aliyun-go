package cbn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AssociateCenBandwidthPackageRequest struct {
	requests.RpcRequest
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	CenId                 string `position:"Query" name:"CenId"`
	CenBandwidthPackageId string `position:"Query" name:"CenBandwidthPackageId"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
}

func (req *AssociateCenBandwidthPackageRequest) Invoke(client *sdk.Client) (resp *AssociateCenBandwidthPackageResponse, err error) {
	req.InitWithApiInfo("Cbn", "2017-09-12", "AssociateCenBandwidthPackage", "cbn", "")
	resp = &AssociateCenBandwidthPackageResponse{}
	err = client.DoAction(req, resp)
	return
}

type AssociateCenBandwidthPackageResponse struct {
	responses.BaseResponse
	RequestId string
}
