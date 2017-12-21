package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListMediaWorkflowExecutionsRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	InputFileURL         string `position:"Query" name:"InputFileURL"`
	NextPageToken        string `position:"Query" name:"NextPageToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	MaximumPageSize      int64  `position:"Query" name:"MaximumPageSize"`
	MediaWorkflowId      string `position:"Query" name:"MediaWorkflowId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	MediaWorkflowName    string `position:"Query" name:"MediaWorkflowName"`
}

func (req *ListMediaWorkflowExecutionsRequest) Invoke(client *sdk.Client) (resp *ListMediaWorkflowExecutionsResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "ListMediaWorkflowExecutions", "", "")
	resp = &ListMediaWorkflowExecutionsResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListMediaWorkflowExecutionsResponse struct {
	responses.BaseResponse
	RequestId                  string
	NextPageToken              string
	MediaWorkflowExecutionList ListMediaWorkflowExecutionsMediaWorkflowExecutionList
}

type ListMediaWorkflowExecutionsMediaWorkflowExecution struct {
	RunId           string
	MediaWorkflowId string
	Name            string
	State           string
	MediaId         string
	CreationTime    string
	ActivityList    ListMediaWorkflowExecutionsActivityList
	Input           ListMediaWorkflowExecutionsInput
}

type ListMediaWorkflowExecutionsActivity struct {
	Name             string
	Type             string
	JobId            string
	State            string
	Code             string
	Message          string
	StartTime        string
	EndTime          string
	MNSMessageResult ListMediaWorkflowExecutionsMNSMessageResult
}

type ListMediaWorkflowExecutionsMNSMessageResult struct {
	MessageId    string
	ErrorMessage string
	ErrorCode    string
}

type ListMediaWorkflowExecutionsInput struct {
	UserData  string
	InputFile ListMediaWorkflowExecutionsInputFile
}

type ListMediaWorkflowExecutionsInputFile struct {
	Bucket   string
	Location string
	Object   string
}

type ListMediaWorkflowExecutionsMediaWorkflowExecutionList []ListMediaWorkflowExecutionsMediaWorkflowExecution

func (list *ListMediaWorkflowExecutionsMediaWorkflowExecutionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListMediaWorkflowExecutionsMediaWorkflowExecution)
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

type ListMediaWorkflowExecutionsActivityList []ListMediaWorkflowExecutionsActivity

func (list *ListMediaWorkflowExecutionsActivityList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListMediaWorkflowExecutionsActivity)
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
