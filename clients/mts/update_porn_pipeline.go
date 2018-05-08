package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type UpdatePornPipelineRequest struct {
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

func (req *UpdatePornPipelineRequest) Invoke(client *sdk.Client) (resp *UpdatePornPipelineResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "UpdatePornPipeline", "mts", "")
	resp = &UpdatePornPipelineResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdatePornPipelineResponse struct {
	responses.BaseResponse
	RequestId common.String
	Pipeline  UpdatePornPipelinePipeline
}

type UpdatePornPipelinePipeline struct {
	Id           common.String
	Name         common.String
	State        common.String
	Priority     common.Integer
	NotifyConfig UpdatePornPipelineNotifyConfig
}

type UpdatePornPipelineNotifyConfig struct {
	Topic common.String
	Queue common.String
}
