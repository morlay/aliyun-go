package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryMediaWorkflowExecutionList", "mts", "")
	resp = &QueryMediaWorkflowExecutionListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryMediaWorkflowExecutionListResponse struct {
	responses.BaseResponse
	RequestId                  common.String
	MediaWorkflowExecutionList QueryMediaWorkflowExecutionListMediaWorkflowExecutionList
	NonExistRunIds             QueryMediaWorkflowExecutionListNonExistRunIdList
}

type QueryMediaWorkflowExecutionListMediaWorkflowExecution struct {
	RunId           common.String
	MediaWorkflowId common.String
	Name            common.String
	State           common.String
	MediaId         common.String
	CreationTime    common.String
	ActivityList    QueryMediaWorkflowExecutionListActivityList
	Input           QueryMediaWorkflowExecutionListInput
}

type QueryMediaWorkflowExecutionListActivity struct {
	Name             common.String
	Type             common.String
	JobId            common.String
	State            common.String
	Code             common.String
	Message          common.String
	StartTime        common.String
	EndTime          common.String
	MNSMessageResult QueryMediaWorkflowExecutionListMNSMessageResult
}

type QueryMediaWorkflowExecutionListMNSMessageResult struct {
	MessageId    common.String
	ErrorMessage common.String
	ErrorCode    common.String
}

type QueryMediaWorkflowExecutionListInput struct {
	UserData  common.String
	InputFile QueryMediaWorkflowExecutionListInputFile
}

type QueryMediaWorkflowExecutionListInputFile struct {
	Bucket   common.String
	Location common.String
	Object   common.String
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

type QueryMediaWorkflowExecutionListNonExistRunIdList []common.String

func (list *QueryMediaWorkflowExecutionListNonExistRunIdList) UnmarshalJSON(data []byte) error {
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
