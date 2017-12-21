package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryTerrorismJobListRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	JobIds               string `position:"Query" name:"JobIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r QueryTerrorismJobListRequest) Invoke(client *sdk.Client) (response *QueryTerrorismJobListResponse, err error) {
	req := struct {
		*requests.RpcRequest
		QueryTerrorismJobListRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryTerrorismJobList", "", "")

	resp := struct {
		*responses.BaseResponse
		QueryTerrorismJobListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.QueryTerrorismJobListResponse

	err = client.DoAction(&req, &resp)
	return
}

type QueryTerrorismJobListResponse struct {
	RequestId        string
	TerrorismJobList QueryTerrorismJobListTerrorismJobList
	NonExistIds      QueryTerrorismJobListNonExistIdList
}

type QueryTerrorismJobListTerrorismJob struct {
	Id                    string
	UserData              string
	PipelineId            string
	State                 string
	Code                  string
	Message               string
	CreationTime          string
	Input                 QueryTerrorismJobListInput
	TerrorismConfig       QueryTerrorismJobListTerrorismConfig
	CensorTerrorismResult QueryTerrorismJobListCensorTerrorismResult
}

type QueryTerrorismJobListInput struct {
	Bucket   string
	Location string
	Object   string
}

type QueryTerrorismJobListTerrorismConfig struct {
	Interval   string
	BizType    string
	OutputFile QueryTerrorismJobListOutputFile
}

type QueryTerrorismJobListOutputFile struct {
	Bucket   string
	Location string
	Object   string
}

type QueryTerrorismJobListCensorTerrorismResult struct {
	Label                string
	Suggestion           string
	MaxScore             string
	AverageScore         string
	TerrorismCounterList QueryTerrorismJobListCounterList
	TerrorismTopList     QueryTerrorismJobListTopList
}

type QueryTerrorismJobListCounter struct {
	Count int
	Label string
}

type QueryTerrorismJobListTop struct {
	Label     string
	Score     string
	Timestamp string
	Index     string
	Object    string
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

type QueryTerrorismJobListNonExistIdList []string

func (list *QueryTerrorismJobListNonExistIdList) UnmarshalJSON(data []byte) error {
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
