package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddAsrPipelineRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	NotifyConfig         string `position:"Query" name:"NotifyConfig"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Priority             int    `position:"Query" name:"Priority"`
}

func (req *AddAsrPipelineRequest) Invoke(client *sdk.Client) (resp *AddAsrPipelineResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "AddAsrPipeline", "mts", "")
	resp = &AddAsrPipelineResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddAsrPipelineResponse struct {
	responses.BaseResponse
	RequestId string
	Pipeline  AddAsrPipelinePipeline
}

type AddAsrPipelinePipeline struct {
	Id           string
	Name         string
	Priority     int
	State        string
	NotifyConfig AddAsrPipelineNotifyConfig
}

type AddAsrPipelineNotifyConfig struct {
	Topic string
	Queue string
}
