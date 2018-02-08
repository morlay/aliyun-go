package aegis

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateInstanceRequest struct {
	requests.RpcRequest
	Duration          int    `position:"Query" name:"Duration"`
	IsAutoRenew       string `position:"Query" name:"IsAutoRenew"`
	ClientToken       string `position:"Query" name:"ClientToken"`
	VmNumber          int    `position:"Query" name:"VmNumber"`
	OwnerId           int64  `position:"Query" name:"OwnerId"`
	VersionCode       int    `position:"Query" name:"VersionCode"`
	PricingCycle      string `position:"Query" name:"PricingCycle"`
	AutoRenewDuration int    `position:"Query" name:"AutoRenewDuration"`
}

func (req *CreateInstanceRequest) Invoke(client *sdk.Client) (resp *CreateInstanceResponse, err error) {
	req.InitWithApiInfo("aegis", "2016-11-11", "CreateInstance", "vipaegis", "")
	resp = &CreateInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateInstanceResponse struct {
	responses.BaseResponse
	OrderId    string
	InstanceId string
	RequestId  string
}
