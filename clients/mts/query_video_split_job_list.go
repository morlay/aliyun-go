package mts

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId   common.String
	JobList     QueryVideoSplitJobListJobList
	NonExistIds QueryVideoSplitJobListNonExistIdList
}

type QueryVideoSplitJobListJob struct {
	Id               common.String
	UserData         common.String
	PipelineId       common.String
	State            common.String
	Code             common.String
	Message          common.String
	CreationTime     common.String
	Input            QueryVideoSplitJobListInput
	VideoSplitResult QueryVideoSplitJobListVideoSplitResult
}

type QueryVideoSplitJobListInput struct {
	Bucket   common.String
	Location common.String
	Object   common.String
}

type QueryVideoSplitJobListVideoSplitResult struct {
	VideoSplitList QueryVideoSplitJobListVideoSplitList
}

type QueryVideoSplitJobListVideoSplit struct {
	StartTime common.String
	EndTime   common.String
	Path      common.String
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

type QueryVideoSplitJobListNonExistIdList []common.String

func (list *QueryVideoSplitJobListNonExistIdList) UnmarshalJSON(data []byte) error {
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
