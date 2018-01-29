package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SubmitPornJobRequest struct {
	requests.RpcRequest
	Input                string `position:"Query" name:"Input"`
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PornConfig           string `position:"Query" name:"PornConfig"`
	PipelineId           string `position:"Query" name:"PipelineId"`
}

func (req *SubmitPornJobRequest) Invoke(client *sdk.Client) (resp *SubmitPornJobResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "SubmitPornJob", "mts", "")
	resp = &SubmitPornJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type SubmitPornJobResponse struct {
	responses.BaseResponse
	RequestId string
	JobId     string
}
