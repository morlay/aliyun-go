package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SubmitAnnotationJobRequest struct {
	Input                string `position:"Query" name:"Input"`
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	AnnotationConfig     string `position:"Query" name:"AnnotationConfig"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PipelineId           string `position:"Query" name:"PipelineId"`
}

func (r SubmitAnnotationJobRequest) Invoke(client *sdk.Client) (response *SubmitAnnotationJobResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SubmitAnnotationJobRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "SubmitAnnotationJob", "", "")

	resp := struct {
		*responses.BaseResponse
		SubmitAnnotationJobResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SubmitAnnotationJobResponse

	err = client.DoAction(&req, &resp)
	return
}

type SubmitAnnotationJobResponse struct {
	RequestId string
	JobId     string
}
