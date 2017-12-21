package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddVideoSummaryPipelineRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	NotifyConfig         string `position:"Query" name:"NotifyConfig"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Priority             int    `position:"Query" name:"Priority"`
}

func (r AddVideoSummaryPipelineRequest) Invoke(client *sdk.Client) (response *AddVideoSummaryPipelineResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddVideoSummaryPipelineRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "AddVideoSummaryPipeline", "", "")

	resp := struct {
		*responses.BaseResponse
		AddVideoSummaryPipelineResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddVideoSummaryPipelineResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddVideoSummaryPipelineResponse struct {
	RequestId string
	Pipeline  AddVideoSummaryPipelinePipeline
}

type AddVideoSummaryPipelinePipeline struct {
	Id           string
	Name         string
	Priority     int
	State        string
	NotifyConfig AddVideoSummaryPipelineNotifyConfig
}

type AddVideoSummaryPipelineNotifyConfig struct {
	Topic string
	Queue string
}
