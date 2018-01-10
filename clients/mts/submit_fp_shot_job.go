package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SubmitFpShotJobRequest struct {
	requests.RpcRequest
	Input                string `position:"Query" name:"Input"`
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	FpShotConfig         string `position:"Query" name:"FpShotConfig"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PipelineId           string `position:"Query" name:"PipelineId"`
}

func (req *SubmitFpShotJobRequest) Invoke(client *sdk.Client) (resp *SubmitFpShotJobResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "SubmitFpShotJob", "", "")
	resp = &SubmitFpShotJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type SubmitFpShotJobResponse struct {
	responses.BaseResponse
	RequestId string
	JobId     string
}
