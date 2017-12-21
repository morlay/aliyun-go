package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReportCoverJobResultRequest struct {
	Result               string `position:"Query" name:"Result"`
	JobId                string `position:"Query" name:"JobId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ReportCoverJobResultRequest) Invoke(client *sdk.Client) (response *ReportCoverJobResultResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ReportCoverJobResultRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "ReportCoverJobResult", "", "")

	resp := struct {
		*responses.BaseResponse
		ReportCoverJobResultResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ReportCoverJobResultResponse

	err = client.DoAction(&req, &resp)
	return
}

type ReportCoverJobResultResponse struct {
	RequestId string
	JobId     string
}
