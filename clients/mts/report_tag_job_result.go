package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReportTagJobResultRequest struct {
	Result               string `position:"Query" name:"Result"`
	JobId                string `position:"Query" name:"JobId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Tag                  string `position:"Query" name:"Tag"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ReportTagJobResultRequest) Invoke(client *sdk.Client) (response *ReportTagJobResultResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ReportTagJobResultRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "ReportTagJobResult", "", "")

	resp := struct {
		*responses.BaseResponse
		ReportTagJobResultResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ReportTagJobResultResponse

	err = client.DoAction(&req, &resp)
	return
}

type ReportTagJobResultResponse struct {
	RequestId string
	JobId     string
}
