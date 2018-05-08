package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ReportFpShotJobResultRequest struct {
	requests.RpcRequest
	Result               string `position:"Query" name:"Result"`
	JobId                string `position:"Query" name:"JobId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Details              string `position:"Query" name:"Details"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ReportFpShotJobResultRequest) Invoke(client *sdk.Client) (resp *ReportFpShotJobResultResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "ReportFpShotJobResult", "mts", "")
	resp = &ReportFpShotJobResultResponse{}
	err = client.DoAction(req, resp)
	return
}

type ReportFpShotJobResultResponse struct {
	responses.BaseResponse
	RequestId common.String
	JobId     common.String
}
