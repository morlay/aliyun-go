package cms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type TaskConfigUnhealthyRequest struct {
	requests.RpcRequest
	TaskIdLists *TaskConfigUnhealthyTaskIdListList `position:"Query" type:"Repeated" name:"TaskIdList"`
}

func (req *TaskConfigUnhealthyRequest) Invoke(client *sdk.Client) (resp *TaskConfigUnhealthyResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "TaskConfigUnhealthy", "cms", "")
	resp = &TaskConfigUnhealthyResponse{}
	err = client.DoAction(req, resp)
	return
}

type TaskConfigUnhealthyResponse struct {
	responses.BaseResponse
	ErrorCode     int
	ErrorMessage  string
	Success       bool
	RequestId     string
	UnhealthyList TaskConfigUnhealthyNodeTaskInstanceList
}

type TaskConfigUnhealthyNodeTaskInstance struct {
	TaskId       int64
	InstanceList TaskConfigUnhealthyInstanceListList
}

type TaskConfigUnhealthyTaskIdListList []int64

func (list *TaskConfigUnhealthyTaskIdListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]int64)
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

type TaskConfigUnhealthyNodeTaskInstanceList []TaskConfigUnhealthyNodeTaskInstance

func (list *TaskConfigUnhealthyNodeTaskInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]TaskConfigUnhealthyNodeTaskInstance)
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

type TaskConfigUnhealthyInstanceListList []string

func (list *TaskConfigUnhealthyInstanceListList) UnmarshalJSON(data []byte) error {
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
