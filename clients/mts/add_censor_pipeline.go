package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddCensorPipelineRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	NotifyConfig         string `position:"Query" name:"NotifyConfig"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Priority             int    `position:"Query" name:"Priority"`
}

func (r AddCensorPipelineRequest) Invoke(client *sdk.Client) (response *AddCensorPipelineResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddCensorPipelineRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "AddCensorPipeline", "", "")

	resp := struct {
		*responses.BaseResponse
		AddCensorPipelineResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddCensorPipelineResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddCensorPipelineResponse struct {
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
