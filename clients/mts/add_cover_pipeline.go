package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddCoverPipelineRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Role                 string `position:"Query" name:"Role"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	NotifyConfig         string `position:"Query" name:"NotifyConfig"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Priority             string `position:"Query" name:"Priority"`
}

func (r AddCoverPipelineRequest) Invoke(client *sdk.Client) (response *AddCoverPipelineResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddCoverPipelineRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "AddCoverPipeline", "", "")

	resp := struct {
		*responses.BaseResponse
		AddCoverPipelineResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddCoverPipelineResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddCoverPipelineResponse struct {
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
