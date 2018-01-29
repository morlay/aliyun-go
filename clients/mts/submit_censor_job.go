package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SubmitCensorJobRequest struct {
	requests.RpcRequest
	Input                string `position:"Query" name:"Input"`
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	CensorConfig         string `position:"Query" name:"CensorConfig"`
	PipelineId           string `position:"Query" name:"PipelineId"`
}

func (req *SubmitCensorJobRequest) Invoke(client *sdk.Client) (resp *SubmitCensorJobResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "SubmitCensorJob", "mts", "")
	resp = &SubmitCensorJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type SubmitCensorJobResponse struct {
	responses.BaseResponse
	RequestId string
	JobId     string
}
