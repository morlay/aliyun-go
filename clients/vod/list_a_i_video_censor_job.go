package vod

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId                   common.String
	AIVideoCensorJobList        ListAIVideoCensorJobAIVideoCensorJobList
	NonExistAIVideoCensorJobIds ListAIVideoCensorJobNonExistAIVideoCensorJobIdList
}

type ListAIVideoCensorJobAIVideoCensorJob struct {
	JobId        common.String
	MediaId      common.String
	Status       common.String
	Code         common.String
	Message      common.String
	CreationTime common.String
	Data         common.String
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

type ListAIVideoCensorJobNonExistAIVideoCensorJobIdList []common.String

func (list *ListAIVideoCensorJobNonExistAIVideoCensorJobIdList) UnmarshalJSON(data []byte) error {
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
