package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryVideoSplitJobListRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	JobIds               string `position:"Query" name:"JobIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *QueryVideoSplitJobListRequest) Invoke(client *sdk.Client) (resp *QueryVideoSplitJobListResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "QueryVideoSplitJobList", "mts", "")
	resp = &QueryVideoSplitJobListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryVideoSplitJobListResponse struct {
	responses.BaseResponse
	RequestId   string
	JobList     QueryVideoSplitJobListJobList
	NonExistIds QueryVideoSplitJobListNonExistIdList
}

type QueryVideoSplitJobListJob struct {
	Id               string
	UserData         string
	PipelineId       string
	State            string
	Code             string
	Message          string
	CreationTime     string
	Input            QueryVideoSplitJobListInput
	VideoSplitResult QueryVideoSplitJobListVideoSplitResult
}

type QueryVideoSplitJobListInput struct {
	Bucket   string
	Location string
	Object   string
}

type QueryVideoSplitJobListVideoSplitResult struct {
	VideoSplitList QueryVideoSplitJobListVideoSplitList
}

type QueryVideoSplitJobListVideoSplit struct {
	StartTime string
	EndTime   string
	Path      string
}

type QueryVideoSplitJobListJobList []QueryVideoSplitJobListJob

func (list *QueryVideoSplitJobListJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryVideoSplitJobListJob)
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

type QueryVideoSplitJobListNonExistIdList []string

func (list *QueryVideoSplitJobListNonExistIdList) UnmarshalJSON(data []byte) error {
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

type QueryVideoSplitJobListVideoSplitList []QueryVideoSplitJobListVideoSplit

func (list *QueryVideoSplitJobListVideoSplitList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryVideoSplitJobListVideoSplit)
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
