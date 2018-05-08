package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryCensorJobListRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	JobIds               string `position:"Query" name:"JobIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *QueryCensorJobListRequest) Invoke(client *sdk.Client) (resp *QueryCensorJobListResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryCensorJobList", "mts", "")
	resp = &QueryCensorJobListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryCensorJobListResponse struct {
	responses.BaseResponse
	RequestId     common.String
	CensorJobList QueryCensorJobListCensorJobList
	NonExistIds   QueryCensorJobListNonExistIdList
}

type QueryCensorJobListCensorJob struct {
	Id                    common.String
	UserData              common.String
	PipelineId            common.String
	State                 common.String
	Code                  common.String
	Message               common.String
	CreationTime          common.String
	Input                 QueryCensorJobListInput
	CensorConfig          QueryCensorJobListCensorConfig
	CensorPornResult      QueryCensorJobListCensorPornResult
	CensorTerrorismResult QueryCensorJobListCensorTerrorismResult
}

type QueryCensorJobListInput struct {
	Bucket   common.String
	Location common.String
	Object   common.String
}

type QueryCensorJobListCensorConfig struct {
	Interval   common.String
	BizType    common.String
	OutputFile QueryCensorJobListOutputFile
}

type QueryCensorJobListOutputFile struct {
	Bucket   common.String
	Location common.String
	Object   common.String
}

type QueryCensorJobListCensorPornResult struct {
	Label           common.String
	Suggestion      common.String
	MaxScore        common.String
	AverageScore    common.String
	PornCounterList QueryCensorJobListCounterList
	PornTopList     QueryCensorJobListTopList
}

type QueryCensorJobListCounter struct {
	Count common.Integer
	Label common.String
}

type QueryCensorJobListTop struct {
	Label     common.String
	Score     common.String
	Timestamp common.String
	Index     common.String
	Object    common.String
}

type QueryCensorJobListCensorTerrorismResult struct {
	Label                common.String
	Suggestion           common.String
	MaxScore             common.String
	AverageScore         common.String
	TerrorismCounterList QueryCensorJobListCounter1List
	TerrorismTopList     QueryCensorJobListTop2List
}

type QueryCensorJobListCounter1 struct {
	Count common.Integer
	Label common.String
}

type QueryCensorJobListTop2 struct {
	Label     common.String
	Score     common.String
	Timestamp common.String
	Index     common.String
	Object    common.String
}

type QueryCensorJobListCensorJobList []QueryCensorJobListCensorJob

func (list *QueryCensorJobListCensorJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryCensorJobListCensorJob)
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

type QueryCensorJobListNonExistIdList []common.String

func (list *QueryCensorJobListNonExistIdList) UnmarshalJSON(data []byte) error {
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

type QueryCensorJobListCounterList []QueryCensorJobListCounter

func (list *QueryCensorJobListCounterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryCensorJobListCounter)
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

type QueryCensorJobListTopList []QueryCensorJobListTop

func (list *QueryCensorJobListTopList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryCensorJobListTop)
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

type QueryCensorJobListCounter1List []QueryCensorJobListCounter1

func (list *QueryCensorJobListCounter1List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryCensorJobListCounter1)
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

type QueryCensorJobListTop2List []QueryCensorJobListTop2

func (list *QueryCensorJobListTop2List) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryCensorJobListTop2)
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
