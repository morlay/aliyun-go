package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type PlayerAuthRequest struct {
	requests.RpcRequest
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

func (req *PlayerAuthRequest) Invoke(client *sdk.Client) (resp *PlayerAuthResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "PlayerAuth", "", "")
	resp = &PlayerAuthResponse{}
	err = client.DoAction(req, resp)
	return
}

type PlayerAuthResponse struct {
	responses.BaseResponse
	RequestId  string
	LogURL     string
	SwitchList PlayerAuth_SwitchList
}

type PlayerAuth_Switch struct {
	State        string
	FunctionId   string
	SwitchId     string
	FunctionName string
}

type PlayerAuth_SwitchList []PlayerAuth_Switch

func (list *PlayerAuth_SwitchList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]PlayerAuth_Switch)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
