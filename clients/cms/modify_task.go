package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyTaskRequest struct {
	requests.RpcRequest
	ReportProject string `position:"Query" name:"ReportProject"`
	TaskType      string `position:"Query" name:"TaskType"`
	Address       string `position:"Query" name:"Address"`
	AlertName     string `position:"Query" name:"AlertName"`
	Ip            string `position:"Query" name:"Ip"`
	AgentGroup    string `position:"Query" name:"AgentGroup"`
	TaskName      string `position:"Query" name:"TaskName"`
	EndTime       string `position:"Query" name:"EndTime"`
	TaskState     string `position:"Query" name:"TaskState"`
	ClientIds     string `position:"Query" name:"ClientIds"`
	AlertInfo     string `position:"Query" name:"AlertInfo"`
	AgentType     string `position:"Query" name:"AgentType"`
	IspCity       string `position:"Query" name:"IspCity"`
	Options       string `position:"Query" name:"Options"`
	Interval      string `position:"Query" name:"Interval"`
	AlertRule     string `position:"Query" name:"AlertRule"`
	TaskId        string `position:"Query" name:"TaskId"`
}

func (req *ModifyTaskRequest) Invoke(client *sdk.Client) (resp *ModifyTaskResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "ModifyTask", "cms", "")
	resp = &ModifyTaskResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyTaskResponse struct {
	responses.BaseResponse
	Code      common.String
	Message   common.String
	Success   common.String
	RequestId common.String
}
