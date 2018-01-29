package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryAsrJobListRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	JobIds               string `position:"Query" name:"JobIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *QueryAsrJobListRequest) Invoke(client *sdk.Client) (resp *QueryAsrJobListResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryAsrJobList", "mts", "")
	resp = &QueryAsrJobListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryAsrJobListResponse struct {
	responses.BaseResponse
	RequestId   string
	JobList     QueryAsrJobListJobList
	NonExistIds QueryAsrJobListNonExistIdList
}

type QueryAsrJobListJob struct {
	Id           string
	UserData     string
	PipelineId   string
	State        string
	Code         string
	Message      string
	CreationTime string
	Input        QueryAsrJobListInput
	AsrConfig    QueryAsrJobListAsrConfig
	AsrResult    QueryAsrJobListAsrResult
}

type QueryAsrJobListInput struct {
	Bucket   string
	Location string
	Object   string
}

type QueryAsrJobListAsrConfig struct {
	Scene string
}

type QueryAsrJobListAsrResult struct {
	Duration    string
	AsrTextList QueryAsrJobListAsrTextList
}

type QueryAsrJobListAsrText struct {
	StartTime  int
	EndTime    string
	ChannelId  string
	SpeechRate string
	Text       string
}

type QueryAsrJobListJobList []QueryAsrJobListJob

func (list *QueryAsrJobListJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryAsrJobListJob)
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

type QueryAsrJobListNonExistIdList []string

func (list *QueryAsrJobListNonExistIdList) UnmarshalJSON(data []byte) error {
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

type QueryAsrJobListAsrTextList []QueryAsrJobListAsrText

func (list *QueryAsrJobListAsrTextList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryAsrJobListAsrText)
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
