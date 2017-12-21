package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryPornJobListRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	JobIds               string `position:"Query" name:"JobIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *QueryPornJobListRequest) Invoke(client *sdk.Client) (resp *QueryPornJobListResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryPornJobList", "", "")
	resp = &QueryPornJobListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryPornJobListResponse struct {
	responses.BaseResponse
	RequestId   string
	PornJobList QueryPornJobListPornJobList
	NonExistIds QueryPornJobListNonExistIdList
}

type QueryPornJobListPornJob struct {
	Id               string
	UserData         string
	PipelineId       string
	State            string
	Code             string
	Message          string
	CreationTime     string
	Input            QueryPornJobListInput
	PornConfig       QueryPornJobListPornConfig
	CensorPornResult QueryPornJobListCensorPornResult
}

type QueryPornJobListInput struct {
	Bucket   string
	Location string
	Object   string
}

type QueryPornJobListPornConfig struct {
	Interval   string
	BizType    string
	OutputFile QueryPornJobListOutputFile
}

type QueryPornJobListOutputFile struct {
	Bucket   string
	Location string
	Object   string
}

type QueryPornJobListCensorPornResult struct {
	Label           string
	Suggestion      string
	MaxScore        string
	AverageScore    string
	PornCounterList QueryPornJobListCounterList
	PornTopList     QueryPornJobListTopList
}

type QueryPornJobListCounter struct {
	Count int
	Label string
}

type QueryPornJobListTop struct {
	Label     string
	Score     string
	Timestamp string
	Index     string
	Object    string
}

type QueryPornJobListPornJobList []QueryPornJobListPornJob

func (list *QueryPornJobListPornJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryPornJobListPornJob)
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

type QueryPornJobListNonExistIdList []string

func (list *QueryPornJobListNonExistIdList) UnmarshalJSON(data []byte) error {
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

type QueryPornJobListCounterList []QueryPornJobListCounter

func (list *QueryPornJobListCounterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryPornJobListCounter)
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

type QueryPornJobListTopList []QueryPornJobListTop

func (list *QueryPornJobListTopList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryPornJobListTop)
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
