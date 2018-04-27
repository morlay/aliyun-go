package cms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type TaskConfigCreateRequest struct {
	requests.RpcRequest
	InstanceLists *TaskConfigCreateInstanceListList `position:"Query" type:"Repeated" name:"InstanceList"`
	JsonData      string                            `position:"Query" name:"JsonData"`
	TaskType      string                            `position:"Query" name:"TaskType"`
	TaskScope     string                            `position:"Query" name:"TaskScope"`
	AlertConfig   string                            `position:"Query" name:"AlertConfig"`
	GroupId       int64                             `position:"Query" name:"GroupId"`
	TaskName      string                            `position:"Query" name:"TaskName"`
	GroupName     string                            `position:"Query" name:"GroupName"`
}

func (req *TaskConfigCreateRequest) Invoke(client *sdk.Client) (resp *TaskConfigCreateResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "TaskConfigCreate", "cms", "")
	resp = &TaskConfigCreateResponse{}
	err = client.DoAction(req, resp)
	return
}

type TaskConfigCreateResponse struct {
	responses.BaseResponse
	ErrorCode    int
	ErrorMessage string
	Success      bool
	RequestId    string
	TaskId       int64
}

type TaskConfigCreateInstanceListList []string

func (list *TaskConfigCreateInstanceListList) UnmarshalJSON(data []byte) error {
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
