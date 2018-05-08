package vod

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	req.InitWithApiInfo("vod", "2017-03-21", "ListAIASRJob", "vod", "")
	resp = &ListAIASRJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListAIASRJobResponse struct {
	responses.BaseResponse
	RequestId           common.String
	AIASRJobList        ListAIASRJobAIASRJobList
	NonExistAIASRJobIds ListAIASRJobNonExistAIASRJobIdList
}

type ListAIASRJobAIASRJob struct {
	JobId        common.String
	MediaId      common.String
	Status       common.String
	Code         common.String
	Message      common.String
	CreationTime common.String
	Data         common.String
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

type ListAIASRJobNonExistAIASRJobIdList []common.String

func (list *ListAIASRJobNonExistAIASRJobIdList) UnmarshalJSON(data []byte) error {
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
