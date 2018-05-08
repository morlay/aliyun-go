package vod

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListAIVideoSummaryJobRequest struct {
	requests.RpcRequest
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	AIVideoSummaryJobIds string `position:"Query" name:"AIVideoSummaryJobIds"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

func (req *ListAIVideoSummaryJobRequest) Invoke(client *sdk.Client) (resp *ListAIVideoSummaryJobResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "ListAIVideoSummaryJob", "vod", "")
	resp = &ListAIVideoSummaryJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListAIVideoSummaryJobResponse struct {
	responses.BaseResponse
	RequestId                    common.String
	AIVideoSummaryJobList        ListAIVideoSummaryJobAIVideoSummaryJobList
	NonExistAIVideoSummaryJobIds ListAIVideoSummaryJobNonExistAIVideoSummaryJobIdList
}

type ListAIVideoSummaryJobAIVideoSummaryJob struct {
	JobId        common.String
	MediaId      common.String
	Status       common.String
	Code         common.String
	Message      common.String
	CreationTime common.String
	Data         common.String
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

type ListAIVideoSummaryJobNonExistAIVideoSummaryJobIdList []common.String

func (list *ListAIVideoSummaryJobNonExistAIVideoSummaryJobIdList) UnmarshalJSON(data []byte) error {
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
