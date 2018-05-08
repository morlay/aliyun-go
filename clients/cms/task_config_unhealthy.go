package cms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	ErrorCode     common.Integer
	ErrorMessage  common.String
	Success       bool
	RequestId     common.String
	UnhealthyList TaskConfigUnhealthyNodeTaskInstanceList
}

type TaskConfigUnhealthyNodeTaskInstance struct {
	TaskId       common.Long
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

type TaskConfigUnhealthyInstanceListList []common.String

func (list *TaskConfigUnhealthyInstanceListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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
