package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddTerrorismPipelineRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	NotifyConfig         string `position:"Query" name:"NotifyConfig"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Priority             int    `position:"Query" name:"Priority"`
}

func (r AddTerrorismPipelineRequest) Invoke(client *sdk.Client) (response *AddTerrorismPipelineResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddTerrorismPipelineRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "AddTerrorismPipeline", "", "")

	resp := struct {
		*responses.BaseResponse
		AddTerrorismPipelineResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddTerrorismPipelineResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddTerrorismPipelineResponse struct {
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
