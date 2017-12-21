package vod

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListAIJobRequest struct {
	requests.RpcRequest
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	JobIds               string `position:"Query" name:"JobIds"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

func (req *ListAIJobRequest) Invoke(client *sdk.Client) (resp *ListAIJobResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "ListAIJob", "", "")
	resp = &ListAIJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListAIJobResponse struct {
	responses.BaseResponse
	RequestId        string
	AIJobList        ListAIJobAIJobList
	NonExistAIJobIds ListAIJobNonExistAIJobIdList
}

type ListAIJobAIJob struct {
	JobId        string
	MediaId      string
	Type         string
	Status       string
	Code         string
	Message      string
	CreationTime string
	CompleteTime string
	Data         string
}

type ListAIJobAIJobList []ListAIJobAIJob

func (list *ListAIJobAIJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAIJobAIJob)
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

type ListAIJobNonExistAIJobIdList []string

func (list *ListAIJobNonExistAIJobIdList) UnmarshalJSON(data []byte) error {
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
