package vod

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListAIASRJobRequest struct {
	requests.RpcRequest
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	AIASRJobIds          string `position:"Query" name:"AIASRJobIds"`
}

func (req *ListAIASRJobRequest) Invoke(client *sdk.Client) (resp *ListAIASRJobResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "ListAIASRJob", "", "")
	resp = &ListAIASRJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListAIASRJobResponse struct {
	responses.BaseResponse
	RequestId           string
	AIASRJobList        ListAIASRJobAIASRJobList
	NonExistAIASRJobIds ListAIASRJobNonExistAIASRJobIdList
}

type ListAIASRJobAIASRJob struct {
	JobId        string
	MediaId      string
	Status       string
	Code         string
	Message      string
	CreationTime string
	Data         string
}

type ListAIASRJobAIASRJobList []ListAIASRJobAIASRJob

func (list *ListAIASRJobAIASRJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAIASRJobAIASRJob)
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

type ListAIASRJobNonExistAIASRJobIdList []string

func (list *ListAIASRJobNonExistAIASRJobIdList) UnmarshalJSON(data []byte) error {
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
