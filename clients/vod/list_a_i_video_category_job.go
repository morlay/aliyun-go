package vod

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListAIVideoCategoryJobRequest struct {
	requests.RpcRequest
	ResourceOwnerId       string `position:"Query" name:"ResourceOwnerId"`
	AIVideoCategoryJobIds string `position:"Query" name:"AIVideoCategoryJobIds"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	OwnerId               string `position:"Query" name:"OwnerId"`
}

func (req *ListAIVideoCategoryJobRequest) Invoke(client *sdk.Client) (resp *ListAIVideoCategoryJobResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "ListAIVideoCategoryJob", "vod", "")
	resp = &ListAIVideoCategoryJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListAIVideoCategoryJobResponse struct {
	responses.BaseResponse
	RequestId                     common.String
	AIVideoCategoryJobList        ListAIVideoCategoryJobAIVideoCategoryJobList
	NonExistAIVideoCategoryJobIds ListAIVideoCategoryJobNonExistAIVideoCategoryJobIdList
}

type ListAIVideoCategoryJobAIVideoCategoryJob struct {
	JobId        common.String
	MediaId      common.String
	Status       common.String
	Code         common.String
	Message      common.String
	CreationTime common.String
	Data         common.String
}

type ListAIVideoCategoryJobAIVideoCategoryJobList []ListAIVideoCategoryJobAIVideoCategoryJob

func (list *ListAIVideoCategoryJobAIVideoCategoryJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAIVideoCategoryJobAIVideoCategoryJob)
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

type ListAIVideoCategoryJobNonExistAIVideoCategoryJobIdList []common.String

func (list *ListAIVideoCategoryJobNonExistAIVideoCategoryJobIdList) UnmarshalJSON(data []byte) error {
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
