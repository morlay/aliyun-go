package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryTerrorismJobListRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	JobIds               string `position:"Query" name:"JobIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *QueryTerrorismJobListRequest) Invoke(client *sdk.Client) (resp *QueryTerrorismJobListResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryTerrorismJobList", "mts", "")
	resp = &QueryTerrorismJobListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryTerrorismJobListResponse struct {
	responses.BaseResponse
	RequestId        common.String
	TerrorismJobList QueryTerrorismJobListTerrorismJobList
	NonExistIds      QueryTerrorismJobListNonExistIdList
}

type QueryTerrorismJobListTerrorismJob struct {
	Id                    common.String
	UserData              common.String
	PipelineId            common.String
	State                 common.String
	Code                  common.String
	Message               common.String
	CreationTime          common.String
	Input                 QueryTerrorismJobListInput
	TerrorismConfig       QueryTerrorismJobListTerrorismConfig
	CensorTerrorismResult QueryTerrorismJobListCensorTerrorismResult
}

type QueryTerrorismJobListInput struct {
	Bucket   common.String
	Location common.String
	Object   common.String
}

type QueryTerrorismJobListTerrorismConfig struct {
	Interval   common.String
	BizType    common.String
	OutputFile QueryTerrorismJobListOutputFile
}

type QueryTerrorismJobListOutputFile struct {
	Bucket   common.String
	Location common.String
	Object   common.String
}

type QueryTerrorismJobListCensorTerrorismResult struct {
	Label                common.String
	Suggestion           common.String
	MaxScore             common.String
	AverageScore         common.String
	TerrorismCounterList QueryTerrorismJobListCounterList
	TerrorismTopList     QueryTerrorismJobListTopList
}

type QueryTerrorismJobListCounter struct {
	Count common.Integer
	Label common.String
}

type QueryTerrorismJobListTop struct {
	Label     common.String
	Score     common.String
	Timestamp common.String
	Index     common.String
	Object    common.String
}

type QueryTerrorismJobListTerrorismJobList []QueryTerrorismJobListTerrorismJob

func (list *QueryTerrorismJobListTerrorismJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryTerrorismJobListTerrorismJob)
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

type QueryTerrorismJobListNonExistIdList []common.String

func (list *QueryTerrorismJobListNonExistIdList) UnmarshalJSON(data []byte) error {
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

type QueryTerrorismJobListCounterList []QueryTerrorismJobListCounter

func (list *QueryTerrorismJobListCounterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryTerrorismJobListCounter)
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

type QueryTerrorismJobListTopList []QueryTerrorismJobListTop

func (list *QueryTerrorismJobListTopList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryTerrorismJobListTop)
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
