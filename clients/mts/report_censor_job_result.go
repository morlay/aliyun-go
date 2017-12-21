package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReportCensorJobResultRequest struct {
	JobId                string `position:"Query" name:"JobId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Label                string `position:"Query" name:"Label"`
	Detail               string `position:"Query" name:"Detail"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ReportCensorJobResultRequest) Invoke(client *sdk.Client) (response *ReportCensorJobResultResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ReportCensorJobResultRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "ReportCensorJobResult", "", "")

	resp := struct {
		*responses.BaseResponse
		ReportCensorJobResultResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ReportCensorJobResultResponse

	err = client.DoAction(&req, &resp)
	return
}

type ReportCensorJobResultResponse struct {
	RequestId string
	JobId     string
}
