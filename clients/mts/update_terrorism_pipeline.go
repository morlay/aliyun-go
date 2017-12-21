package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	req.InitWithApiInfo("Mts", "2014-06-18", "UpdateTerrorismPipeline", "", "")
	resp = &UpdateTerrorismPipelineResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateTerrorismPipelineResponse struct {
	responses.BaseResponse
	RequestId string
	Pipeline  UpdateTerrorismPipelinePipeline
}

type UpdateTerrorismPipelinePipeline struct {
	Id           string
	Name         string
	State        string
	Priority     int
	NotifyConfig UpdateTerrorismPipelineNotifyConfig
}

type UpdateTerrorismPipelineNotifyConfig struct {
	Topic string
	Queue string
}
