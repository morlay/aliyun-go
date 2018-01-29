package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SubmitTerrorismJobRequest struct {
	requests.RpcRequest
	Input                string `position:"Query" name:"Input"`
	UserData             string `position:"Query" name:"UserData"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PipelineId           string `position:"Query" name:"PipelineId"`
	TerrorismConfig      string `position:"Query" name:"TerrorismConfig"`
}

func (req *SubmitTerrorismJobRequest) Invoke(client *sdk.Client) (resp *SubmitTerrorismJobResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "SubmitTerrorismJob", "mts", "")
	resp = &SubmitTerrorismJobResponse{}
	err = client.DoAction(req, resp)
	return
}

type SubmitTerrorismJobResponse struct {
	responses.BaseResponse
	RequestId string
	JobId     string
}
