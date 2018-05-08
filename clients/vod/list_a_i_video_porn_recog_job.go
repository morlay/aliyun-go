package vod

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListAIVideoPornRecogJobRequest struct {
	requests.RpcRequest
	ResourceOwnerId        string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount   string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount           string `position:"Query" name:"OwnerAccount"`
	AIVideoPornRecogJobIds string `position:"Query" name:"AIVideoPornRecogJobIds"`
	OwnerId                string `position:"Query" name:"OwnerId"`
}

func (req *ListAIVideoPornRecogJobRequest) Invoke(client *sdk.Client) (resp *ListAIVideoPornRecogJobResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "ListAIVideoPornRecogJob", "vod", "")
	resp = &ListAIVideoPornRecogJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListAIVideoPornRecogJobResponse struct {
	responses.BaseResponse
	RequestId               common.String
	AIVideoPornRecogJobList ListAIVideoPornRecogJobAIVideoPornRecogJobList
	NonExistPornRecogJobIds ListAIVideoPornRecogJobNonExistPornRecogJobIdList
}

type ListAIVideoPornRecogJobAIVideoPornRecogJob struct {
	JobId        common.String
	MediaId      common.String
	Status       common.String
	Code         common.String
	Message      common.String
	CreationTime common.String
	Data         common.String
}

type ListAIVideoPornRecogJobAIVideoPornRecogJobList []ListAIVideoPornRecogJobAIVideoPornRecogJob

func (list *ListAIVideoPornRecogJobAIVideoPornRecogJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAIVideoPornRecogJobAIVideoPornRecogJob)
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

type ListAIVideoPornRecogJobNonExistPornRecogJobIdList []common.String

func (list *ListAIVideoPornRecogJobNonExistPornRecogJobIdList) UnmarshalJSON(data []byte) error {
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
