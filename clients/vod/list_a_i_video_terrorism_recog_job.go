package vod

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListAIVideoTerrorismRecogJobRequest struct {
	requests.RpcRequest
	ResourceOwnerId             string `position:"Query" name:"ResourceOwnerId"`
	AIVideoTerrorismRecogJobIds string `position:"Query" name:"AIVideoTerrorismRecogJobIds"`
	ResourceOwnerAccount        string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                string `position:"Query" name:"OwnerAccount"`
	OwnerId                     string `position:"Query" name:"OwnerId"`
}

func (req *ListAIVideoTerrorismRecogJobRequest) Invoke(client *sdk.Client) (resp *ListAIVideoTerrorismRecogJobResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "ListAIVideoTerrorismRecogJob", "vod", "")
	resp = &ListAIVideoTerrorismRecogJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListAIVideoTerrorismRecogJobResponse struct {
	responses.BaseResponse
	RequestId                    common.String
	AIVideoTerrorismRecogJobList ListAIVideoTerrorismRecogJobAIVideoTerrorismRecogJobList
	NonExistTerrorismRecogJobIds ListAIVideoTerrorismRecogJobNonExistTerrorismRecogJobIdList
}

type ListAIVideoTerrorismRecogJobAIVideoTerrorismRecogJob struct {
	JobId        common.String
	MediaId      common.String
	Status       common.String
	Code         common.String
	Message      common.String
	CreationTime common.String
	Data         common.String
}

type ListAIVideoTerrorismRecogJobAIVideoTerrorismRecogJobList []ListAIVideoTerrorismRecogJobAIVideoTerrorismRecogJob

func (list *ListAIVideoTerrorismRecogJobAIVideoTerrorismRecogJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAIVideoTerrorismRecogJobAIVideoTerrorismRecogJob)
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

type ListAIVideoTerrorismRecogJobNonExistTerrorismRecogJobIdList []common.String

func (list *ListAIVideoTerrorismRecogJobNonExistTerrorismRecogJobIdList) UnmarshalJSON(data []byte) error {
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
