package cbn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateCenBandwidthPackageRequest struct {
	requests.RpcRequest
	GeographicRegionBId        string `position:"Query" name:"GeographicRegionBId"`
	ResourceOwnerId            int64  `position:"Query" name:"ResourceOwnerId"`
	Period                     int    `position:"Query" name:"Period"`
	GeographicRegionAId        string `position:"Query" name:"GeographicRegionAId"`
	AutoPay                    string `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount       string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken                string `position:"Query" name:"ClientToken"`
	Bandwidth                  int    `position:"Query" name:"Bandwidth"`
	OwnerAccount               string `position:"Query" name:"OwnerAccount"`
	Description                string `position:"Query" name:"Description"`
	OwnerId                    int64  `position:"Query" name:"OwnerId"`
	BandwidthPackageChargeType string `position:"Query" name:"BandwidthPackageChargeType"`
	Name                       string `position:"Query" name:"Name"`
	PricingCycle               string `position:"Query" name:"PricingCycle"`
}

func (req *CreateCenBandwidthPackageRequest) Invoke(client *sdk.Client) (resp *CreateCenBandwidthPackageResponse, err error) {
	req.InitWithApiInfo("Cbn", "2017-09-12", "CreateCenBandwidthPackage", "cbn", "")
	resp = &CreateCenBandwidthPackageResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateCenBandwidthPackageResponse struct {
	responses.BaseResponse
	RequestId                  common.String
	CenBandwidthPackageId      common.String
	CenBandwidthPackageOrderId common.String
}
