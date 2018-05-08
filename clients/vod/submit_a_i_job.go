package vod

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SubmitAIJobRequest struct {
	requests.RpcRequest
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	Types                string `position:"Query" name:"Types"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	MediaId              string `position:"Query" name:"MediaId"`
	Config               string `position:"Query" name:"Config"`
}

func (req *SubmitAIJobRequest) Invoke(client *sdk.Client) (resp *SubmitAIJobResponse, err error) {
	req.InitWithApiInfo("vod", "2017-03-21", "SubmitAIJob", "vod", "")
	resp = &SubmitAIJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type SubmitAIJobResponse struct {
	responses.BaseResponse
	RequestId common.String
	AIJobList SubmitAIJobAIJobList
}

type SubmitAIJobAIJob struct {
	JobId        common.String
	Type         common.String
	MediaId      common.String
	Status       common.String
	Code         common.String
	Message      common.String
	CreationTime common.String
	Data         common.String
}

type SubmitAIJobAIJobList []SubmitAIJobAIJob

func (list *SubmitAIJobAIJobList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]SubmitAIJobAIJob)
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
