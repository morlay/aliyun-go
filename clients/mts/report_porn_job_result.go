package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReportPornJobResultRequest struct {
	JobId                string `position:"Query" name:"JobId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Label                string `position:"Query" name:"Label"`
	Detail               string `position:"Query" name:"Detail"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ReportPornJobResultRequest) Invoke(client *sdk.Client) (response *ReportPornJobResultResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ReportPornJobResultRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "ReportPornJobResult", "", "")

	resp := struct {
		*responses.BaseResponse
		ReportPornJobResultResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ReportPornJobResultResponse

	err = client.DoAction(&req, &resp)
	return
}

type ReportPornJobResultResponse struct {
	RequestId string
	JobId     string
}
