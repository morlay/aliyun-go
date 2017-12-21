package vod

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListAIVideoSummaryJobRequest struct {
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	AIVideoSummaryJobIds string `position:"Query" name:"AIVideoSummaryJobIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

func (r ListAIVideoSummaryJobRequest) Invoke(client *sdk.Client) (response *ListAIVideoSummaryJobResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ListAIVideoSummaryJobRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("vod", "2017-03-21", "ListAIVideoSummaryJob", "", "")

	resp := struct {
		*responses.BaseResponse
		ListAIVideoSummaryJobResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ListAIVideoSummaryJobResponse

	err = client.DoAction(&req, &resp)
	return
}

type ListAIVideoSummaryJobResponse struct {
	RequestId                    string
	AIVideoSummaryJobList        ListAIVideoSummaryJobAIVideoSummaryJobList
	NonExistAIVideoSummaryJobIds ListAIVideoSummaryJobNonExistAIVideoSummaryJobIdList
}

type ListAIVideoSummaryJobAIVideoSummaryJob struct {
	JobId        string
	MediaId      string
	Status       string
	Code         string
	Message      string
	CreationTime string
	Data         string
}

type ListAIVideoSummaryJobAIVideoSummaryJobList []ListAIVideoSummaryJobAIVideoSummaryJob

func (list *ListAIVideoSummaryJobAIVideoSummaryJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAIVideoSummaryJobAIVideoSummaryJob)
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

type ListAIVideoSummaryJobNonExistAIVideoSummaryJobIdList []string

func (list *ListAIVideoSummaryJobNonExistAIVideoSummaryJobIdList) UnmarshalJSON(data []byte) error {
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
