package vod

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListAIVideoCoverJobRequest struct {
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	AIVideoCoverJobIds   string `position:"Query" name:"AIVideoCoverJobIds"`
	OwnerId              string `position:"Query" name:"OwnerId"`
}

func (r ListAIVideoCoverJobRequest) Invoke(client *sdk.Client) (response *ListAIVideoCoverJobResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ListAIVideoCoverJobRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("vod", "2017-03-21", "ListAIVideoCoverJob", "", "")

	resp := struct {
		*responses.BaseResponse
		ListAIVideoCoverJobResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ListAIVideoCoverJobResponse

	err = client.DoAction(&req, &resp)
	return
}

type ListAIVideoCoverJobResponse struct {
	RequestId                  string
	AIVideoCoverJobList        ListAIVideoCoverJobAIVideoCoverJobList
	NonExistAIVideoCoverJobIds ListAIVideoCoverJobNonExistAIVideoCoverJobIdList
}

type ListAIVideoCoverJobAIVideoCoverJob struct {
	JobId        string
	MediaId      string
	Status       string
	Code         string
	Message      string
	CreationTime string
	Data         string
}

type ListAIVideoCoverJobAIVideoCoverJobList []ListAIVideoCoverJobAIVideoCoverJob

func (list *ListAIVideoCoverJobAIVideoCoverJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListAIVideoCoverJobAIVideoCoverJob)
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

type ListAIVideoCoverJobNonExistAIVideoCoverJobIdList []string

func (list *ListAIVideoCoverJobNonExistAIVideoCoverJobIdList) UnmarshalJSON(data []byte) error {
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
