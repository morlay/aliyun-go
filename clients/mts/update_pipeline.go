package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type UpdatePipelineRequest struct {
	requests.RpcRequest
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

func (req *UpdatePipelineRequest) Invoke(client *sdk.Client) (resp *UpdatePipelineResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "UpdatePipeline", "mts", "")
	resp = &UpdatePipelineResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdatePipelineResponse struct {
	responses.BaseResponse
	RequestId common.String
	Pipeline  UpdatePipelinePipeline
}

type UpdatePipelinePipeline struct {
	Id           common.String
	Name         common.String
	State        common.String
	Speed        common.String
	Role         common.String
	NotifyConfig UpdatePipelineNotifyConfig
}

type UpdatePipelineNotifyConfig struct {
	Topic     common.String
	QueueName common.String
}
