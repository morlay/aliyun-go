package vod

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	req.InitWithApiInfo("vod", "2017-03-21", "ListAIJob", "vod", "")
	resp = &ListAIJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListAIJobResponse struct {
	responses.BaseResponse
	RequestId        common.String
	AIJobList        ListAIJobAIJobList
	NonExistAIJobIds ListAIJobNonExistAIJobIdList
}

type ListAIJobAIJob struct {
	JobId        common.String
	MediaId      common.String
	Type         common.String
	Status       common.String
	Code         common.String
	Message      common.String
	CreationTime common.String
	CompleteTime common.String
	Data         common.String
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

type ListAIJobNonExistAIJobIdList []common.String

func (list *ListAIJobNonExistAIJobIdList) UnmarshalJSON(data []byte) error {
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
