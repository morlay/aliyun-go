package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReportCoverJobResultRequest struct {
	requests.RpcRequest
	Result               string `position:"Query" name:"Result"`
	JobId                string `position:"Query" name:"JobId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ReportCoverJobResultRequest) Invoke(client *sdk.Client) (resp *ReportCoverJobResultResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "ReportCoverJobResult", "", "")
	resp = &ReportCoverJobResultResponse{}
	err = client.DoAction(req, resp)
	return
}

type ReportCoverJobResultResponse struct {
	responses.BaseResponse
	RequestId string
	JobId     string
}
