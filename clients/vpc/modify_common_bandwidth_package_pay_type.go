package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyCommonBandwidthPackagePayTypeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	BandwidthPackageId   string `position:"Query" name:"BandwidthPackageId"`
	AutoPay              string `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth            string `position:"Query" name:"Bandwidth"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Duration             int    `position:"Query" name:"Duration"`
	KbpsBandwidth        string `position:"Query" name:"KbpsBandwidth"`
	ResourceUid          int64  `position:"Query" name:"ResourceUid"`
	ResourceBid          string `position:"Query" name:"ResourceBid"`
	PayType              string `position:"Query" name:"PayType"`
	PricingCycle         string `position:"Query" name:"PricingCycle"`
}

func (r ModifyCommonBandwidthPackagePayTypeRequest) Invoke(client *sdk.Client) (response *ModifyCommonBandwidthPackagePayTypeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyCommonBandwidthPackagePayTypeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "ModifyCommonBandwidthPackagePayType", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		ModifyCommonBandwidthPackagePayTypeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyCommonBandwidthPackagePayTypeResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyCommonBandwidthPackagePayTypeResponse struct {
	RequestId string
	OrderId   int64
	Code      string
	Message   string
}
