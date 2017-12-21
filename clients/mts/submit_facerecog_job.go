package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SubmitFacerecogJobRequest struct {
	Input                string `position:"Query" name:"Input"`
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	FacerecogConfig      string `position:"Query" name:"FacerecogConfig"`
	PipelineId           string `position:"Query" name:"PipelineId"`
}

func (r SubmitFacerecogJobRequest) Invoke(client *sdk.Client) (response *SubmitFacerecogJobResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SubmitFacerecogJobRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "SubmitFacerecogJob", "", "")

	resp := struct {
		*responses.BaseResponse
		SubmitFacerecogJobResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SubmitFacerecogJobResponse

	err = client.DoAction(&req, &resp)
	return
}

type SubmitFacerecogJobResponse struct {
	RequestId string
	JobId     string
}
