package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdatePornPipelineRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	State                string `position:"Query" name:"State"`
	NotifyConfig         string `position:"Query" name:"NotifyConfig"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Priority             int    `position:"Query" name:"Priority"`
	PipelineId           string `position:"Query" name:"PipelineId"`
}

func (r UpdatePornPipelineRequest) Invoke(client *sdk.Client) (response *UpdatePornPipelineResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UpdatePornPipelineRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "UpdatePornPipeline", "", "")

	resp := struct {
		*responses.BaseResponse
		UpdatePornPipelineResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UpdatePornPipelineResponse

	err = client.DoAction(&req, &resp)
	return
}

type UpdatePornPipelineResponse struct {
	RequestId string
	Pipeline  UpdatePornPipelinePipeline
}

type UpdatePornPipelinePipeline struct {
	Id           string
	Name         string
	State        string
	Priority     int
	NotifyConfig UpdatePornPipelineNotifyConfig
}

type UpdatePornPipelineNotifyConfig struct {
	Topic string
	Queue string
}
