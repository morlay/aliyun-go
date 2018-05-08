package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SubmitVideoSummaryJobRequest struct {
	requests.RpcRequest
	Input                string `position:"Query" name:"Input"`
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	VideoSummaryConfig   string `position:"Query" name:"VideoSummaryConfig"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PipelineId           string `position:"Query" name:"PipelineId"`
}

func (req *SubmitVideoSummaryJobRequest) Invoke(client *sdk.Client) (resp *SubmitVideoSummaryJobResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "SubmitVideoSummaryJob", "mts", "")
	resp = &SubmitVideoSummaryJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type SubmitVideoSummaryJobResponse struct {
	responses.BaseResponse
	RequestId common.String
	JobId     common.String
}
