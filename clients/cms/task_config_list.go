package cms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	ErrorCode    common.Integer
	ErrorMessage common.String
	Success      bool
	RequestId    common.String
	PageNumber   common.Integer
	PageSize     common.Integer
	PageTotal    common.Integer
	Total        common.Integer
	TaskList     TaskConfigListNodeTaskConfigList
}

type TaskConfigListNodeTaskConfig struct {
	Id           common.Long
	TaskName     common.String
	TaskType     common.String
	TaskScope    common.String
	Disabled     bool
	GroupId      common.Long
	GroupName    common.String
	JsonData     common.String
	AlertConfig  common.String
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

type TaskConfigListInstanceListList []common.String

func (list *TaskConfigListInstanceListList) UnmarshalJSON(data []byte) error {
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
