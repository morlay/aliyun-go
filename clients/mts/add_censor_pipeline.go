package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddCensorPipelineRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	NotifyConfig         string `position:"Query" name:"NotifyConfig"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Priority             int    `position:"Query" name:"Priority"`
}

func (req *AddCensorPipelineRequest) Invoke(client *sdk.Client) (resp *AddCensorPipelineResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "AddCensorPipeline", "", "")
	resp = &AddCensorPipelineResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddCensorPipelineResponse struct {
	responses.BaseResponse
	RequestId string
	Pipeline  AddCensorPipelinePipeline
}

type AddCensorPipelinePipeline struct {
	Id           string
	Name         string
	Priority     int
	State        string
	NotifyConfig AddCensorPipelineNotifyConfig
}

type AddCensorPipelineNotifyConfig struct {
	Topic string
	Queue string
}
