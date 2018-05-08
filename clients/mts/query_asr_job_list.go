package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId   common.String
	JobList     QueryAsrJobListJobList
	NonExistIds QueryAsrJobListNonExistIdList
}

type QueryAsrJobListJob struct {
	Id           common.String
	UserData     common.String
	PipelineId   common.String
	State        common.String
	Code         common.String
	Message      common.String
	CreationTime common.String
	Input        QueryAsrJobListInput
	AsrConfig    QueryAsrJobListAsrConfig
	AsrResult    QueryAsrJobListAsrResult
}

type QueryAsrJobListInput struct {
	Bucket   common.String
	Location common.String
	Object   common.String
}

type QueryAsrJobListAsrConfig struct {
	Scene common.String
}

type QueryAsrJobListAsrResult struct {
	Duration    common.String
	AsrTextList QueryAsrJobListAsrTextList
}

type QueryAsrJobListAsrText struct {
	StartTime  common.Integer
	EndTime    common.String
	ChannelId  common.String
	SpeechRate common.String
	Text       common.String
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

type QueryAsrJobListNonExistIdList []common.String

func (list *QueryAsrJobListNonExistIdList) UnmarshalJSON(data []byte) error {
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
