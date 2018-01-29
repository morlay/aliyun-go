package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddCoverPipelineRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Role                 string `position:"Query" name:"Role"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	NotifyConfig         string `position:"Query" name:"NotifyConfig"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Priority             string `position:"Query" name:"Priority"`
}

func (req *AddCoverPipelineRequest) Invoke(client *sdk.Client) (resp *AddCoverPipelineResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "AddCoverPipeline", "mts", "")
	resp = &AddCoverPipelineResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddCoverPipelineResponse struct {
	responses.BaseResponse
	RequestId string
	Pipeline  AddCoverPipelinePipeline
}

type AddCoverPipelinePipeline struct {
	Id           string
	Name         string
	Priority     string
	State        string
	Role         string
	NotifyConfig AddCoverPipelineNotifyConfig
}

type AddCoverPipelineNotifyConfig struct {
	Topic string
	Queue string
}
