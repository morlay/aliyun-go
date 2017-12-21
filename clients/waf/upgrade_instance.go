package waf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpgradeInstanceRequest struct {
	requests.RpcRequest
	OwnerId          int64  `position:"Query" name:"OwnerId"`
	ClientToken      string `position:"Query" name:"ClientToken"`
	InstanceId       string `position:"Query" name:"InstanceId"`
	PackageCode      string `position:"Query" name:"PackageCode"`
	ExtDomainPackage int    `position:"Query" name:"ExtDomainPackage"`
	ExtBandwidth     int    `position:"Query" name:"ExtBandwidth"`
}

func (req *UpgradeInstanceRequest) Invoke(client *sdk.Client) (resp *UpgradeInstanceResponse, err error) {
	req.InitWithApiInfo("waf-openapi", "2016-11-11", "UpgradeInstance", "", "")
	resp = &UpgradeInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpgradeInstanceResponse struct {
	responses.BaseResponse
	OrderId   string
	RequestId string
}
