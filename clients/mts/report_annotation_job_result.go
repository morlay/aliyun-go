package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReportAnnotationJobResultRequest struct {
	Annotation           string `position:"Query" name:"Annotation"`
	JobId                string `position:"Query" name:"JobId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Details              string `position:"Query" name:"Details"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ReportAnnotationJobResultRequest) Invoke(client *sdk.Client) (response *ReportAnnotationJobResultResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ReportAnnotationJobResultRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "ReportAnnotationJobResult", "", "")

	resp := struct {
		*responses.BaseResponse
		ReportAnnotationJobResultResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ReportAnnotationJobResultResponse

	err = client.DoAction(&req, &resp)
	return
}

type ReportAnnotationJobResultResponse struct {
	RequestId string
	JobId     string
}
