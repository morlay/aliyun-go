package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RegisterMediaDetailScenarioRequest struct {
	requests.RpcRequest
	JobId                string `position:"Query" name:"JobId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Scenario             string `position:"Query" name:"Scenario"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *RegisterMediaDetailScenarioRequest) Invoke(client *sdk.Client) (resp *RegisterMediaDetailScenarioResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "RegisterMediaDetailScenario", "", "")
	resp = &RegisterMediaDetailScenarioResponse{}
	err = client.DoAction(req, resp)
	return
}

type RegisterMediaDetailScenarioResponse struct {
	responses.BaseResponse
	RequestId  string
	ScenarioId string
}
