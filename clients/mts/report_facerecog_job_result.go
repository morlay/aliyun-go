package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReportFacerecogJobResultRequest struct {
	JobId                string `position:"Query" name:"JobId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Facerecog            string `position:"Query" name:"Facerecog"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Details              string `position:"Query" name:"Details"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ReportFacerecogJobResultRequest) Invoke(client *sdk.Client) (response *ReportFacerecogJobResultResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ReportFacerecogJobResultRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "ReportFacerecogJobResult", "", "")

	resp := struct {
		*responses.BaseResponse
		ReportFacerecogJobResultResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ReportFacerecogJobResultResponse

	err = client.DoAction(&req, &resp)
	return
}

type ReportFacerecogJobResultResponse struct {
	RequestId string
	JobId     string
}
