package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryMediaWorkflowExecutionListRequest struct {
	requests.RpcRequest
	RunIds               string `position:"Query" name:"RunIds"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *QueryMediaWorkflowExecutionListRequest) Invoke(client *sdk.Client) (resp *QueryMediaWorkflowExecutionListResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryMediaWorkflowExecutionList", "", "")
	resp = &QueryMediaWorkflowExecutionListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryMediaWorkflowExecutionListResponse struct {
	responses.BaseResponse
	RequestId                  string
	MediaWorkflowExecutionList QueryMediaWorkflowExecutionListMediaWorkflowExecutionList
	NonExistRunIds             QueryMediaWorkflowExecutionListNonExistRunIdList
}

type QueryMediaWorkflowExecutionListMediaWorkflowExecution struct {
	RunId           string
	MediaWorkflowId string
	Name            string
	State           string
	MediaId         string
	CreationTime    string
	ActivityList    QueryMediaWorkflowExecutionListActivityList
	Input           QueryMediaWorkflowExecutionListInput
}

type QueryMediaWorkflowExecutionListActivity struct {
	Name             string
	Type             string
	JobId            string
	State            string
	Code             string
	Message          string
	StartTime        string
	EndTime          string
	MNSMessageResult QueryMediaWorkflowExecutionListMNSMessageResult
}

type QueryMediaWorkflowExecutionListMNSMessageResult struct {
	MessageId    string
	ErrorMessage string
	ErrorCode    string
}

type QueryMediaWorkflowExecutionListInput struct {
	UserData  string
	InputFile QueryMediaWorkflowExecutionListInputFile
}

type QueryMediaWorkflowExecutionListInputFile struct {
	Bucket   string
	Location string
	Object   string
}

type QueryMediaWorkflowExecutionListMediaWorkflowExecutionList []QueryMediaWorkflowExecutionListMediaWorkflowExecution

func (list *QueryMediaWorkflowExecutionListMediaWorkflowExecutionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaWorkflowExecutionListMediaWorkflowExecution)
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

type QueryMediaWorkflowExecutionListNonExistRunIdList []string

func (list *QueryMediaWorkflowExecutionListNonExistRunIdList) UnmarshalJSON(data []byte) error {
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

type QueryMediaWorkflowExecutionListActivityList []QueryMediaWorkflowExecutionListActivity

func (list *QueryMediaWorkflowExecutionListActivityList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryMediaWorkflowExecutionListActivity)
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
