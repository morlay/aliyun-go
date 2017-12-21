package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateCensorPipelineRequest struct {
	requests.RpcRequest
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

func (req *UpdateCensorPipelineRequest) Invoke(client *sdk.Client) (resp *UpdateCensorPipelineResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "UpdateCensorPipeline", "", "")
	resp = &UpdateCensorPipelineResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateCensorPipelineResponse struct {
	responses.BaseResponse
	RequestId string
	Pipeline  UpdateCensorPipelinePipeline
}

type UpdateCensorPipelinePipeline struct {
	Id           string
	Name         string
	State        string
	Priority     int
	NotifyConfig UpdateCensorPipelineNotifyConfig
}

type UpdateCensorPipelineNotifyConfig struct {
	Topic string
	Queue string
}
