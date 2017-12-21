package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateAsrPipelineRequest struct {
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

func (r UpdateAsrPipelineRequest) Invoke(client *sdk.Client) (response *UpdateAsrPipelineResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UpdateAsrPipelineRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "UpdateAsrPipeline", "", "")

	resp := struct {
		*responses.BaseResponse
		UpdateAsrPipelineResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UpdateAsrPipelineResponse

	err = client.DoAction(&req, &resp)
	return
}

type UpdateAsrPipelineResponse struct {
	RequestId string
	Pipeline  UpdateAsrPipelinePipeline
}

type UpdateAsrPipelinePipeline struct {
	Id           string
	Name         string
	State        string
	Priority     int
	NotifyConfig UpdateAsrPipelineNotifyConfig
}

type UpdateAsrPipelineNotifyConfig struct {
	Topic     string
	QueueName string
}
