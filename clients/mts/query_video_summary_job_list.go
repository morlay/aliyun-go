package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryVideoSummaryJobListRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	JobIds               string `position:"Query" name:"JobIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *QueryVideoSummaryJobListRequest) Invoke(client *sdk.Client) (resp *QueryVideoSummaryJobListResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryVideoSummaryJobList", "", "")
	resp = &QueryVideoSummaryJobListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryVideoSummaryJobListResponse struct {
	responses.BaseResponse
	RequestId   string
	JobList     QueryVideoSummaryJobListJobList
	NonExistIds QueryVideoSummaryJobListNonExistIdList
}

type QueryVideoSummaryJobListJob struct {
	Id                 string
	UserData           string
	PipelineId         string
	State              string
	Code               string
	Message            string
	CreationTime       string
	Input              QueryVideoSummaryJobListInput
	VideoSummaryResult QueryVideoSummaryJobListVideoSummaryResult
}

type QueryVideoSummaryJobListInput struct {
	Bucket   string
	Location string
	Object   string
}

type QueryVideoSummaryJobListVideoSummaryResult struct {
	VideoSummaryList QueryVideoSummaryJobListVideoSummaryList
}

type QueryVideoSummaryJobListVideoSummary struct {
	StartTime string
	EndTime   string
}

type QueryVideoSummaryJobListJobList []QueryVideoSummaryJobListJob

func (list *QueryVideoSummaryJobListJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryVideoSummaryJobListJob)
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

type QueryVideoSummaryJobListNonExistIdList []string

func (list *QueryVideoSummaryJobListNonExistIdList) UnmarshalJSON(data []byte) error {
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

type QueryVideoSummaryJobListVideoSummaryList []QueryVideoSummaryJobListVideoSummary

func (list *QueryVideoSummaryJobListVideoSummaryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryVideoSummaryJobListVideoSummary)
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
