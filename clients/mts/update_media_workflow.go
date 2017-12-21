package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateMediaWorkflowRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Topology             string `position:"Query" name:"Topology"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	MediaWorkflowId      string `position:"Query" name:"MediaWorkflowId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r UpdateMediaWorkflowRequest) Invoke(client *sdk.Client) (response *UpdateMediaWorkflowResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UpdateMediaWorkflowRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "UpdateMediaWorkflow", "", "")

	resp := struct {
		*responses.BaseResponse
		UpdateMediaWorkflowResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UpdateMediaWorkflowResponse

	err = client.DoAction(&req, &resp)
	return
}

type UpdateMediaWorkflowResponse struct {
	RequestId     string
	MediaWorkflow UpdateMediaWorkflowMediaWorkflow
}

type UpdateMediaWorkflowMediaWorkflow struct {
	MediaWorkflowId string
	Name            string
	Topology        string
	State           string
	CreationTime    string
}
