package waf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpgradeInstanceRequest struct {
	OwnerId          int64  `position:"Query" name:"OwnerId"`
	ClientToken      string `position:"Query" name:"ClientToken"`
	InstanceId       string `position:"Query" name:"InstanceId"`
	PackageCode      string `position:"Query" name:"PackageCode"`
	ExtDomainPackage int    `position:"Query" name:"ExtDomainPackage"`
	ExtBandwidth     int    `position:"Query" name:"ExtBandwidth"`
}

func (r UpgradeInstanceRequest) Invoke(client *sdk.Client) (response *UpgradeInstanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UpgradeInstanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("waf-openapi", "2016-11-11", "UpgradeInstance", "", "")

	resp := struct {
		*responses.BaseResponse
		UpgradeInstanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UpgradeInstanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type UpgradeInstanceResponse struct {
	OrderId   string
	RequestId string
}
