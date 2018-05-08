package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type UpdateVideoSummaryPipelineRequest struct {
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

func (req *UpdateVideoSummaryPipelineRequest) Invoke(client *sdk.Client) (resp *UpdateVideoSummaryPipelineResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "UpdateVideoSummaryPipeline", "mts", "")
	resp = &UpdateVideoSummaryPipelineResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateVideoSummaryPipelineResponse struct {
	responses.BaseResponse
	RequestId common.String
	Pipeline  UpdateVideoSummaryPipelinePipeline
}

type UpdateVideoSummaryPipelinePipeline struct {
	Id           common.String
	Name         common.String
	State        common.String
	Priority     common.Integer
	NotifyConfig UpdateVideoSummaryPipelineNotifyConfig
}

type UpdateVideoSummaryPipelineNotifyConfig struct {
	Topic     common.String
	QueueName common.String
}
