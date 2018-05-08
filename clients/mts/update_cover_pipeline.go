package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type UpdateCoverPipelineRequest struct {
	requests.RpcRequest
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

func (req *UpdateCoverPipelineRequest) Invoke(client *sdk.Client) (resp *UpdateCoverPipelineResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "UpdateCoverPipeline", "mts", "")
	resp = &UpdateCoverPipelineResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateCoverPipelineResponse struct {
	responses.BaseResponse
	RequestId common.String
	Pipeline  UpdateCoverPipelinePipeline
}

type UpdateCoverPipelinePipeline struct {
	Id           common.String
	Name         common.String
	State        common.String
	Priority     common.Integer
	Role         common.String
	NotifyConfig UpdateCoverPipelineNotifyConfig
}

type UpdateCoverPipelineNotifyConfig struct {
	Topic common.String
	Queue common.String
}
