package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeactivateMediaWorkflowRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	MediaWorkflowId      string `position:"Query" name:"MediaWorkflowId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DeactivateMediaWorkflowRequest) Invoke(client *sdk.Client) (response *DeactivateMediaWorkflowResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeactivateMediaWorkflowRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "DeactivateMediaWorkflow", "", "")

	resp := struct {
		*responses.BaseResponse
		DeactivateMediaWorkflowResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeactivateMediaWorkflowResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeactivateMediaWorkflowResponse struct {
	RequestId     string
	MediaWorkflow DeactivateMediaWorkflowMediaWorkflow
}

type DeactivateMediaWorkflowMediaWorkflow struct {
	MediaWorkflowId string
	Name            string
	Topology        string
	State           string
	CreationTime    string
}
