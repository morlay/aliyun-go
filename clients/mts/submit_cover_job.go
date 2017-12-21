package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SubmitCoverJobRequest struct {
	requests.RpcRequest
	Input                string `position:"Query" name:"Input"`
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	CoverConfig          string `position:"Query" name:"CoverConfig"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PipelineId           string `position:"Query" name:"PipelineId"`
}

func (req *SubmitCoverJobRequest) Invoke(client *sdk.Client) (resp *SubmitCoverJobResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "SubmitCoverJob", "", "")
	resp = &SubmitCoverJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type SubmitCoverJobResponse struct {
	responses.BaseResponse
	RequestId string
	JobId     string
}
