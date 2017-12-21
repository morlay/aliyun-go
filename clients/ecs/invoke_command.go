package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type InvokeCommandRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64                        `position:"Query" name:"ResourceOwnerId"`
	CommandId            string                       `position:"Query" name:"CommandId"`
	Frequency            string                       `position:"Query" name:"Frequency"`
	Timed                string                       `position:"Query" name:"Timed"`
	ResourceOwnerAccount string                       `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                       `position:"Query" name:"OwnerAccount"`
	OwnerId              int64                        `position:"Query" name:"OwnerId"`
	InstanceIds          *InvokeCommandInstanceIdList `position:"Query" type:"Repeated" name:"InstanceId"`
}

func (req *InvokeCommandRequest) Invoke(client *sdk.Client) (resp *InvokeCommandResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "InvokeCommand", "ecs", "")
	resp = &InvokeCommandResponse{}
	err = client.DoAction(req, resp)
	return
}

type InvokeCommandResponse struct {
	responses.BaseResponse
	RequestId string
	InvokeId  string
}

type InvokeCommandInstanceIdList []string

func (list *InvokeCommandInstanceIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
