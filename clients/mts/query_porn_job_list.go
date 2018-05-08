package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryPornJobList", "mts", "")
	resp = &QueryPornJobListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryPornJobListResponse struct {
	responses.BaseResponse
	RequestId   common.String
	PornJobList QueryPornJobListPornJobList
	NonExistIds QueryPornJobListNonExistIdList
}

type QueryPornJobListPornJob struct {
	Id               common.String
	UserData         common.String
	PipelineId       common.String
	State            common.String
	Code             common.String
	Message          common.String
	CreationTime     common.String
	Input            QueryPornJobListInput
	PornConfig       QueryPornJobListPornConfig
	CensorPornResult QueryPornJobListCensorPornResult
}

type QueryPornJobListInput struct {
	Bucket   common.String
	Location common.String
	Object   common.String
}

type QueryPornJobListPornConfig struct {
	Interval   common.String
	BizType    common.String
	OutputFile QueryPornJobListOutputFile
}

type QueryPornJobListOutputFile struct {
	Bucket   common.String
	Location common.String
	Object   common.String
}

type QueryPornJobListCensorPornResult struct {
	Label           common.String
	Suggestion      common.String
	MaxScore        common.String
	AverageScore    common.String
	PornCounterList QueryPornJobListCounterList
	PornTopList     QueryPornJobListTopList
}

type QueryPornJobListCounter struct {
	Count common.Integer
	Label common.String
}

type QueryPornJobListTop struct {
	Label     common.String
	Score     common.String
	Timestamp common.String
	Index     common.String
	Object    common.String
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

type QueryPornJobListNonExistIdList []common.String

func (list *QueryPornJobListNonExistIdList) UnmarshalJSON(data []byte) error {
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
