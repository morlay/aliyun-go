package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AddVideoSummaryPipelineRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	NotifyConfig         string `position:"Query" name:"NotifyConfig"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Priority             int    `position:"Query" name:"Priority"`
}

func (req *AddVideoSummaryPipelineRequest) Invoke(client *sdk.Client) (resp *AddVideoSummaryPipelineResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "AddVideoSummaryPipeline", "mts", "")
	resp = &AddVideoSummaryPipelineResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddVideoSummaryPipelineResponse struct {
	responses.BaseResponse
	RequestId common.String
	Pipeline  AddVideoSummaryPipelinePipeline
}

type AddVideoSummaryPipelinePipeline struct {
	Id           common.String
	Name         common.String
	Priority     common.Integer
	State        common.String
	NotifyConfig AddVideoSummaryPipelineNotifyConfig
}

type AddVideoSummaryPipelineNotifyConfig struct {
	Topic common.String
	Queue common.String
}
