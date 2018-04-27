package cms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type TaskConfigListRequest struct {
	requests.RpcRequest
	GroupId    int64  `position:"Query" name:"GroupId"`
	PageSize   int    `position:"Query" name:"PageSize"`
	TaskName   string `position:"Query" name:"TaskName"`
	Id         int64  `position:"Query" name:"Id"`
	PageNumber int    `position:"Query" name:"PageNumber"`
}

func (req *TaskConfigListRequest) Invoke(client *sdk.Client) (resp *TaskConfigListResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "TaskConfigList", "cms", "")
	resp = &TaskConfigListResponse{}
	err = client.DoAction(req, resp)
	return
}

type TaskConfigListResponse struct {
	responses.BaseResponse
	ErrorCode    int
	ErrorMessage string
	Success      bool
	RequestId    string
	PageNumber   int
	PageSize     int
	PageTotal    int
	Total        int
	TaskList     TaskConfigListNodeTaskConfigList
}

type TaskConfigListNodeTaskConfig struct {
	Id           int64
	TaskName     string
	TaskType     string
	TaskScope    string
	Disabled     bool
	GroupId      int64
	GroupName    string
	JsonData     string
	AlertConfig  string
	InstanceList TaskConfigListInstanceListList
}

type TaskConfigListNodeTaskConfigList []TaskConfigListNodeTaskConfig

func (list *TaskConfigListNodeTaskConfigList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]TaskConfigListNodeTaskConfig)
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

type TaskConfigListInstanceListList []string

func (list *TaskConfigListInstanceListList) UnmarshalJSON(data []byte) error {
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
