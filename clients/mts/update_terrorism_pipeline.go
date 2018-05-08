package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type UpdateTerrorismPipelineRequest struct {
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

func (req *UpdateTerrorismPipelineRequest) Invoke(client *sdk.Client) (resp *UpdateTerrorismPipelineResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "UpdateTerrorismPipeline", "mts", "")
	resp = &UpdateTerrorismPipelineResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateTerrorismPipelineResponse struct {
	responses.BaseResponse
	RequestId common.String
	Pipeline  UpdateTerrorismPipelinePipeline
}

type UpdateTerrorismPipelinePipeline struct {
	Id           common.String
	Name         common.String
	State        common.String
	Priority     common.Integer
	NotifyConfig UpdateTerrorismPipelineNotifyConfig
}

type UpdateTerrorismPipelineNotifyConfig struct {
	Topic common.String
	Queue common.String
}
