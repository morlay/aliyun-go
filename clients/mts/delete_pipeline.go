package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeletePipelineRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PipelineId           string `position:"Query" name:"PipelineId"`
}

func (r DeletePipelineRequest) Invoke(client *sdk.Client) (response *DeletePipelineResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeletePipelineRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "DeletePipeline", "", "")

	resp := struct {
		*responses.BaseResponse
		DeletePipelineResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeletePipelineResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeletePipelineResponse struct {
	RequestId  string
	PipelineId string
}
