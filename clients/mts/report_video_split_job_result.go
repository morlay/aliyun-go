package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ReportVideoSplitJobResultRequest struct {
	requests.RpcRequest
	Result               string `position:"Query" name:"Result"`
	JobId                string `position:"Query" name:"JobId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Details              string `position:"Query" name:"Details"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ReportVideoSplitJobResultRequest) Invoke(client *sdk.Client) (resp *ReportVideoSplitJobResultResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "ReportVideoSplitJobResult", "mts", "")
	resp = &ReportVideoSplitJobResultResponse{}
	err = client.DoAction(req, resp)
	return
}

type ReportVideoSplitJobResultResponse struct {
	responses.BaseResponse
	RequestId common.String
	JobId     common.String
}
