package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdatePipelineRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Role                 string `position:"Query" name:"Role"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	State                string `position:"Query" name:"State"`
	NotifyConfig         string `position:"Query" name:"NotifyConfig"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PipelineId           string `position:"Query" name:"PipelineId"`
}

func (r UpdatePipelineRequest) Invoke(client *sdk.Client) (response *UpdatePipelineResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UpdatePipelineRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "UpdatePipeline", "", "")

	resp := struct {
		*responses.BaseResponse
		UpdatePipelineResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UpdatePipelineResponse

	err = client.DoAction(&req, &resp)
	return
}

type UpdatePipelineResponse struct {
	RequestId string
	Pipeline  UpdatePipelinePipeline
}

type UpdatePipelinePipeline struct {
	Id           string
	Name         string
	State        string
	Speed        string
	Role         string
	NotifyConfig UpdatePipelineNotifyConfig
}

type UpdatePipelineNotifyConfig struct {
	Topic     string
	QueueName string
}
