package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReportTerrorismJobResultRequest struct {
	requests.RpcRequest
	JobId                string `position:"Query" name:"JobId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Label                string `position:"Query" name:"Label"`
	Detail               string `position:"Query" name:"Detail"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ReportTerrorismJobResultRequest) Invoke(client *sdk.Client) (resp *ReportTerrorismJobResultResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "ReportTerrorismJobResult", "mts", "")
	resp = &ReportTerrorismJobResultResponse{}
	err = client.DoAction(req, resp)
	return
}

type ReportTerrorismJobResultResponse struct {
	responses.BaseResponse
	RequestId string
	JobId     string
}
