package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateCoverPipelineRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Role                 string `position:"Query" name:"Role"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	State                string `position:"Query" name:"State"`
	NotifyConfig         string `position:"Query" name:"NotifyConfig"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Priority             int    `position:"Query" name:"Priority"`
	PipelineId           string `position:"Query" name:"PipelineId"`
}

func (r UpdateCoverPipelineRequest) Invoke(client *sdk.Client) (response *UpdateCoverPipelineResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UpdateCoverPipelineRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "UpdateCoverPipeline", "", "")

	resp := struct {
		*responses.BaseResponse
		UpdateCoverPipelineResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UpdateCoverPipelineResponse

	err = client.DoAction(&req, &resp)
	return
}

type UpdateCoverPipelineResponse struct {
	RequestId string
	Pipeline  UpdateCoverPipelinePipeline
}

type UpdateCoverPipelinePipeline struct {
	Id           string
	Name         string
	State        string
	Priority     int
	Role         string
	NotifyConfig UpdateCoverPipelineNotifyConfig
}

type UpdateCoverPipelineNotifyConfig struct {
	Topic string
	Queue string
}
