package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddAsrPipelineRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	NotifyConfig         string `position:"Query" name:"NotifyConfig"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Priority             int    `position:"Query" name:"Priority"`
}

func (r AddAsrPipelineRequest) Invoke(client *sdk.Client) (response *AddAsrPipelineResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddAsrPipelineRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "AddAsrPipeline", "", "")

	resp := struct {
		*responses.BaseResponse
		AddAsrPipelineResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddAsrPipelineResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddAsrPipelineResponse struct {
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
