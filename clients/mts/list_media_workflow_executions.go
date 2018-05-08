package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	req.InitWithApiInfo("Mts", "2014-06-18", "ListMediaWorkflowExecutions", "mts", "")
	resp = &ListMediaWorkflowExecutionsResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListMediaWorkflowExecutionsResponse struct {
	responses.BaseResponse
	RequestId                  common.String
	NextPageToken              common.String
	MediaWorkflowExecutionList ListMediaWorkflowExecutionsMediaWorkflowExecutionList
}

type ListMediaWorkflowExecutionsMediaWorkflowExecution struct {
	RunId           common.String
	MediaWorkflowId common.String
	Name            common.String
	State           common.String
	MediaId         common.String
	CreationTime    common.String
	ActivityList    ListMediaWorkflowExecutionsActivityList
	Input           ListMediaWorkflowExecutionsInput
}

type ListMediaWorkflowExecutionsActivity struct {
	Name             common.String
	Type             common.String
	JobId            common.String
	State            common.String
	Code             common.String
	Message          common.String
	StartTime        common.String
	EndTime          common.String
	MNSMessageResult ListMediaWorkflowExecutionsMNSMessageResult
}

type ListMediaWorkflowExecutionsMNSMessageResult struct {
	MessageId    common.String
	ErrorMessage common.String
	ErrorCode    common.String
}

type ListMediaWorkflowExecutionsInput struct {
	UserData  common.String
	InputFile ListMediaWorkflowExecutionsInputFile
}

type ListMediaWorkflowExecutionsInputFile struct {
	Bucket   common.String
	Location common.String
	Object   common.String
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
