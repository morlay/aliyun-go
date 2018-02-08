package aegis

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RenewInstanceRequest struct {
	requests.RpcRequest
	Duration     int    `position:"Query" name:"Duration"`
	InstanceId   string `position:"Query" name:"InstanceId"`
	ClientToken  string `position:"Query" name:"ClientToken"`
	VmNumber     string `position:"Query" name:"VmNumber"`
	OwnerId      int64  `position:"Query" name:"OwnerId"`
	PricingCycle string `position:"Query" name:"PricingCycle"`
}

func (req *RenewInstanceRequest) Invoke(client *sdk.Client) (resp *RenewInstanceResponse, err error) {
	req.InitWithApiInfo("aegis", "2016-11-11", "RenewInstance", "vipaegis", "")
	resp = &RenewInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type RenewInstanceResponse struct {
	responses.BaseResponse
	OrderId   string
	RequestId string
}
