package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeletePipelineRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PipelineId           string `position:"Query" name:"PipelineId"`
}

func (req *DeletePipelineRequest) Invoke(client *sdk.Client) (resp *DeletePipelineResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "DeletePipeline", "", "")
	resp = &DeletePipelineResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeletePipelineResponse struct {
	responses.BaseResponse
	RequestId  string
	PipelineId string
}
