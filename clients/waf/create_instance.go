package waf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateInstanceRequest struct {
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

func (r CreateInstanceRequest) Invoke(client *sdk.Client) (response *CreateInstanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateInstanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("waf-openapi", "2016-11-11", "CreateInstance", "", "")

	resp := struct {
		*responses.BaseResponse
		CreateInstanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateInstanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateInstanceResponse struct {
	OrderId    string
	InstanceId string
	RequestId  string
}
