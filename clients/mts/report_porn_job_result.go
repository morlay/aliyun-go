package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ReportPornJobResultRequest struct {
	requests.RpcRequest
	JobId                string `position:"Query" name:"JobId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Label                string `position:"Query" name:"Label"`
	Detail               string `position:"Query" name:"Detail"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ReportPornJobResultRequest) Invoke(client *sdk.Client) (resp *ReportPornJobResultResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "ReportPornJobResult", "mts", "")
	resp = &ReportPornJobResultResponse{}
	err = client.DoAction(req, resp)
	return
}

type ReportPornJobResultResponse struct {
	responses.BaseResponse
	RequestId common.String
	JobId     common.String
}
