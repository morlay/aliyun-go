package ccc

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type StartJobRequest struct {
	requests.RpcRequest
	JobJson              string                     `position:"Query" name:"JobJson"`
	CallingNumbers       *StartJobCallingNumberList `position:"Query" type:"Repeated" name:"CallingNumber"`
	InstanceId           string                     `position:"Query" name:"InstanceId"`
	GroupId              string                     `position:"Query" name:"GroupId"`
	SelfHostedCallCenter string                     `position:"Query" name:"SelfHostedCallCenter"`
	ScenarioId           string                     `position:"Query" name:"ScenarioId"`
}

func (req *StartJobRequest) Invoke(client *sdk.Client) (resp *StartJobResponse, err error) {
	req.InitWithApiInfo("CCC", "2017-07-05", "StartJob", "ccc", "")
	resp = &StartJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type StartJobResponse struct {
	responses.BaseResponse
	RequestId      string
	Success        bool
	Code           string
	Message        string
	HttpStatusCode int
	TaskIds        StartJobKeyValuePairList
}

type StartJobKeyValuePair struct {
	Key   string
	Value string
}

type StartJobCallingNumberList []string

func (list *StartJobCallingNumberList) UnmarshalJSON(data []byte) error {
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

type StartJobKeyValuePairList []StartJobKeyValuePair

func (list *StartJobKeyValuePairList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]StartJobKeyValuePair)
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
