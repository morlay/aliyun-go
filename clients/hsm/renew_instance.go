package hsm

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RenewInstanceRequest struct {
	requests.RpcRequest
	Period          int    `position:"Query" name:"Period"`
	PeriodUnit      string `position:"Query" name:"PeriodUnit"`
	ResourceOwnerId string `position:"Query" name:"ResourceOwnerId"`
	InstanceId      string `position:"Query" name:"InstanceId"`
	ClientToken     string `position:"Query" name:"ClientToken"`
}

func (req *RenewInstanceRequest) Invoke(client *sdk.Client) (resp *RenewInstanceResponse, err error) {
	req.InitWithApiInfo("hsm", "2018-01-11", "RenewInstance", "hsm", "")
	resp = &RenewInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type RenewInstanceResponse struct {
	responses.BaseResponse
	RequestId string
}
