package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReportMediaDetailJobResultRequest struct {
	requests.RpcRequest
	JobId                string `position:"Query" name:"JobId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Tag                  string `position:"Query" name:"Tag"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Results              string `position:"Query" name:"Results"`
}

func (req *ReportMediaDetailJobResultRequest) Invoke(client *sdk.Client) (resp *ReportMediaDetailJobResultResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "ReportMediaDetailJobResult", "mts", "")
	resp = &ReportMediaDetailJobResultResponse{}
	err = client.DoAction(req, resp)
	return
}

type ReportMediaDetailJobResultResponse struct {
	responses.BaseResponse
	RequestId string
	JobId     string
}
