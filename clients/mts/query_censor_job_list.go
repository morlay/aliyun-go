package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryCensorJobListRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	JobIds               string `position:"Query" name:"JobIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r QueryCensorJobListRequest) Invoke(client *sdk.Client) (response *QueryCensorJobListResponse, err error) {
	req := struct {
		*requests.RpcRequest
		QueryCensorJobListRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryCensorJobList", "", "")

	resp := struct {
		*responses.BaseResponse
		QueryCensorJobListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.QueryCensorJobListResponse

	err = client.DoAction(&req, &resp)
	return
}

type QueryCensorJobListResponse struct {
	RequestId     string
	CensorJobList QueryCensorJobListCensorJobList
	NonExistIds   QueryCensorJobListNonExistIdList
}

type QueryCensorJobListCensorJob struct {
	Id                    string
	UserData              string
	PipelineId            string
	State                 string
	Code                  string
	Message               string
	CreationTime          string
	Input                 QueryCensorJobListInput
	CensorConfig          QueryCensorJobListCensorConfig
	CensorPornResult      QueryCensorJobListCensorPornResult
	CensorTerrorismResult QueryCensorJobListCensorTerrorismResult
}

type QueryCensorJobListInput struct {
	Bucket   string
	Location string
	Object   string
}

type QueryCensorJobListCensorConfig struct {
	Interval   string
	BizType    string
	OutputFile QueryCensorJobListOutputFile
}

type QueryCensorJobListOutputFile struct {
	Bucket   string
	Location string
	Object   string
}

type QueryCensorJobListCensorPornResult struct {
	Label           string
	Suggestion      string
	MaxScore        string
	AverageScore    string
	PornCounterList QueryCensorJobListCounterList
	PornTopList     QueryCensorJobListTopList
}

type QueryCensorJobListCounter struct {
	Count int
	Label string
}

type QueryCensorJobListTop struct {
	Label     string
	Score     string
	Timestamp string
	Index     string
	Object    string
}

type QueryCensorJobListCensorTerrorismResult struct {
	Label                string
	Suggestion           string
	MaxScore             string
	AverageScore         string
	TerrorismCounterList QueryCensorJobListCounter1List
	TerrorismTopList     QueryCensorJobListTop2List
}

type QueryCensorJobListCounter1 struct {
	Count int
	Label string
}

type QueryCensorJobListTop2 struct {
	Label     string
	Score     string
	Timestamp string
	Index     string
	Object    string
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

type QueryCensorJobListNonExistIdList []string

func (list *QueryCensorJobListNonExistIdList) UnmarshalJSON(data []byte) error {
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
