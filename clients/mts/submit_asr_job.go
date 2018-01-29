package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SubmitAsrJobRequest struct {
	requests.RpcRequest
	Input                string `position:"Query" name:"Input"`
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	AsrConfig            string `position:"Query" name:"AsrConfig"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PipelineId           string `position:"Query" name:"PipelineId"`
}

func (req *SubmitAsrJobRequest) Invoke(client *sdk.Client) (resp *SubmitAsrJobResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "SubmitAsrJob", "mts", "")
	resp = &SubmitAsrJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type SubmitAsrJobResponse struct {
	responses.BaseResponse
	RequestId string
	JobId     string
}
