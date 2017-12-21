package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddMediaWorkflowRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Topology             string `position:"Query" name:"Topology"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r AddMediaWorkflowRequest) Invoke(client *sdk.Client) (response *AddMediaWorkflowResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddMediaWorkflowRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "AddMediaWorkflow", "", "")

	resp := struct {
		*responses.BaseResponse
		AddMediaWorkflowResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddMediaWorkflowResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddMediaWorkflowResponse struct {
	RequestId     string
	MediaWorkflow AddMediaWorkflowMediaWorkflow
}

type AddMediaWorkflowMediaWorkflow struct {
	MediaWorkflowId string
	Name            string
	Topology        string
	State           string
	CreationTime    string
}
