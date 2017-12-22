package vod

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListAIVideoCensorJobRequest struct {
	requests.RpcRequest
	AIVideoCensorJobIds  string `position:"Query" name:"AIVideoCensorJobIds"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

func (req *ListAIVideoCensorJobRequest) Invoke(client *sdk.Client) (resp *ListAIVideoCensorJobResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "ListAIVideoCensorJob", "vod", "")
	resp = &ListAIVideoCensorJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListAIVideoCensorJobResponse struct {
	responses.BaseResponse
	RequestId                   string
	AIVideoCensorJobList        ListAIVideoCensorJobAIVideoCensorJobList
	NonExistAIVideoCensorJobIds ListAIVideoCensorJobNonExistAIVideoCensorJobIdList
}

type ListAIVideoCensorJobAIVideoCensorJob struct {
	JobId        string
	MediaId      string
	Status       string
	Code         string
	Message      string
	CreationTime string
	Data         string
}

type ListAIVideoCensorJobAIVideoCensorJobList []ListAIVideoCensorJobAIVideoCensorJob

func (list *ListAIVideoCensorJobAIVideoCensorJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAIVideoCensorJobAIVideoCensorJob)
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

type ListAIVideoCensorJobNonExistAIVideoCensorJobIdList []string

func (list *ListAIVideoCensorJobNonExistAIVideoCensorJobIdList) UnmarshalJSON(data []byte) error {
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
