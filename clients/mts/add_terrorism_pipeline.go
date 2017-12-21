package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddTerrorismPipelineRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	NotifyConfig         string `position:"Query" name:"NotifyConfig"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Priority             int    `position:"Query" name:"Priority"`
}

func (req *AddTerrorismPipelineRequest) Invoke(client *sdk.Client) (resp *AddTerrorismPipelineResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "AddTerrorismPipeline", "", "")
	resp = &AddTerrorismPipelineResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddTerrorismPipelineResponse struct {
	responses.BaseResponse
	RequestId string
	Pipeline  AddTerrorismPipelinePipeline
}

type AddTerrorismPipelinePipeline struct {
	Id           string
	Name         string
	Priority     int
	State        string
	NotifyConfig AddTerrorismPipelineNotifyConfig
}

type AddTerrorismPipelineNotifyConfig struct {
	Topic string
	Queue string
}
