package vod

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SubmitAIJobRequest struct {
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	Types                string `position:"Query" name:"Types"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	MediaId              string `position:"Query" name:"MediaId"`
	Config               string `position:"Query" name:"Config"`
}

func (r SubmitAIJobRequest) Invoke(client *sdk.Client) (response *SubmitAIJobResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SubmitAIJobRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("vod", "2017-03-21", "SubmitAIJob", "", "")

	resp := struct {
		*responses.BaseResponse
		SubmitAIJobResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SubmitAIJobResponse

	err = client.DoAction(&req, &resp)
	return
}

type SubmitAIJobResponse struct {
	RequestId string
	AIJobList SubmitAIJobAIJobList
}

type SubmitAIJobAIJob struct {
	JobId        string
	Type         string
	MediaId      string
	Status       string
	Code         string
	Message      string
	CreationTime string
	Data         string
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
