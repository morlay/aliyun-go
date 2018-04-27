package cms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type TaskConfigModifyRequest struct {
	requests.RpcRequest
	InstanceLists *TaskConfigModifyInstanceListList `position:"Query" type:"Repeated" name:"InstanceList"`
	JsonData      string                            `position:"Query" name:"JsonData"`
	TaskType      string                            `position:"Query" name:"TaskType"`
	TaskScope     string                            `position:"Query" name:"TaskScope"`
	AlertConfig   string                            `position:"Query" name:"AlertConfig"`
	GroupId       int64                             `position:"Query" name:"GroupId"`
	TaskName      string                            `position:"Query" name:"TaskName"`
	Id            int64                             `position:"Query" name:"Id"`
	GroupName     string                            `position:"Query" name:"GroupName"`
}

func (req *TaskConfigModifyRequest) Invoke(client *sdk.Client) (resp *TaskConfigModifyResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "TaskConfigModify", "cms", "")
	resp = &TaskConfigModifyResponse{}
	err = client.DoAction(req, resp)
	return
}

type TaskConfigModifyResponse struct {
	responses.BaseResponse
	ErrorCode    int
	ErrorMessage string
	Success      bool
	RequestId    string
}

type TaskConfigModifyInstanceListList []string

func (list *TaskConfigModifyInstanceListList) UnmarshalJSON(data []byte) error {
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
