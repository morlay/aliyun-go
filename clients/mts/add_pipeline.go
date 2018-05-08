package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AddPipelineRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Role                 string `position:"Query" name:"Role"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	NotifyConfig         string `position:"Query" name:"NotifyConfig"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	SpeedLevel           int64  `position:"Query" name:"SpeedLevel"`
	Speed                string `position:"Query" name:"Speed"`
}

func (req *AddPipelineRequest) Invoke(client *sdk.Client) (resp *AddPipelineResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "AddPipeline", "mts", "")
	resp = &AddPipelineResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddPipelineResponse struct {
	responses.BaseResponse
	RequestId common.String
	Pipeline  AddPipelinePipeline
}

type AddPipelinePipeline struct {
	Id           common.String
	Name         common.String
	State        common.String
	Speed        common.String
	SpeedLevel   common.Long
	Role         common.String
	NotifyConfig AddPipelineNotifyConfig
}

type AddPipelineNotifyConfig struct {
	Topic     common.String
	QueueName common.String
}
