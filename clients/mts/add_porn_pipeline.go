package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddPornPipelineRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	NotifyConfig         string `position:"Query" name:"NotifyConfig"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Priority             int    `position:"Query" name:"Priority"`
}

func (req *AddPornPipelineRequest) Invoke(client *sdk.Client) (resp *AddPornPipelineResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "AddPornPipeline", "", "")
	resp = &AddPornPipelineResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddPornPipelineResponse struct {
	responses.BaseResponse
	RequestId string
	Pipeline  AddPornPipelinePipeline
}

type AddPornPipelinePipeline struct {
	Id           string
	Name         string
	Priority     int
	State        string
	NotifyConfig AddPornPipelineNotifyConfig
}

type AddPornPipelineNotifyConfig struct {
	Topic string
	Queue string
}
