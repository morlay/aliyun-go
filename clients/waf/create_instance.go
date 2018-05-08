package waf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateInstanceRequest struct {
	requests.RpcRequest
	OwnerId           int64  `position:"Query" name:"OwnerId"`
	ClientToken       string `position:"Query" name:"ClientToken"`
	Duration          int    `position:"Query" name:"Duration"`
	PricingCycle      string `position:"Query" name:"PricingCycle"`
	PackageCode       string `position:"Query" name:"PackageCode"`
	ExtDomainPackage  int    `position:"Query" name:"ExtDomainPackage"`
	ExtBandwidth      int    `position:"Query" name:"ExtBandwidth"`
	IsAutoRenew       string `position:"Query" name:"IsAutoRenew"`
	AutoRenewDuration int    `position:"Query" name:"AutoRenewDuration"`
}

func (req *CreateInstanceRequest) Invoke(client *sdk.Client) (resp *CreateInstanceResponse, err error) {
	req.InitWithApiInfo("waf-openapi", "2016-11-11", "CreateInstance", "", "")
	resp = &CreateInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateInstanceResponse struct {
	responses.BaseResponse
	OrderId    common.String
	InstanceId common.String
	RequestId  common.String
}
