package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryVideoSummaryJobList", "mts", "")
	resp = &QueryVideoSummaryJobListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryVideoSummaryJobListResponse struct {
	responses.BaseResponse
	RequestId   common.String
	JobList     QueryVideoSummaryJobListJobList
	NonExistIds QueryVideoSummaryJobListNonExistIdList
}

type QueryVideoSummaryJobListJob struct {
	Id                 common.String
	UserData           common.String
	PipelineId         common.String
	State              common.String
	Code               common.String
	Message            common.String
	CreationTime       common.String
	Input              QueryVideoSummaryJobListInput
	VideoSummaryResult QueryVideoSummaryJobListVideoSummaryResult
}

type QueryVideoSummaryJobListInput struct {
	Bucket   common.String
	Location common.String
	Object   common.String
}

type QueryVideoSummaryJobListVideoSummaryResult struct {
	VideoSummaryList QueryVideoSummaryJobListVideoSummaryList
}

type QueryVideoSummaryJobListVideoSummary struct {
	StartTime common.String
	EndTime   common.String
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

type QueryVideoSummaryJobListNonExistIdList []common.String

func (list *QueryVideoSummaryJobListNonExistIdList) UnmarshalJSON(data []byte) error {
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
