package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SubmitVideoSplitJobRequest struct {
	requests.RpcRequest
	Input                string `position:"Query" name:"Input"`
	VideoSplitConfig     string `position:"Query" name:"VideoSplitConfig"`
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PipelineId           string `position:"Query" name:"PipelineId"`
}

func (req *SubmitVideoSplitJobRequest) Invoke(client *sdk.Client) (resp *SubmitVideoSplitJobResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "SubmitVideoSplitJob", "mts", "")
	resp = &SubmitVideoSplitJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type SubmitVideoSplitJobResponse struct {
	responses.BaseResponse
	RequestId string
	JobId     string
}
