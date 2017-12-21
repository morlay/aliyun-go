package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyCommonBandwidthPackagePayTypeRequest struct {
	requests.RpcRequest
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

func (req *ModifyCommonBandwidthPackagePayTypeRequest) Invoke(client *sdk.Client) (resp *ModifyCommonBandwidthPackagePayTypeResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "ModifyCommonBandwidthPackagePayType", "vpc", "")
	resp = &ModifyCommonBandwidthPackagePayTypeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyCommonBandwidthPackagePayTypeResponse struct {
	responses.BaseResponse
	RequestId string
	OrderId   int64
	Code      string
	Message   string
}
