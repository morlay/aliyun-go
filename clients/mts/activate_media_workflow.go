package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ActivateMediaWorkflowRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	MediaWorkflowId      string `position:"Query" name:"MediaWorkflowId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ActivateMediaWorkflowRequest) Invoke(client *sdk.Client) (response *ActivateMediaWorkflowResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ActivateMediaWorkflowRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "ActivateMediaWorkflow", "", "")

	resp := struct {
		*responses.BaseResponse
		ActivateMediaWorkflowResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ActivateMediaWorkflowResponse

	err = client.DoAction(&req, &resp)
	return
}

type ActivateMediaWorkflowResponse struct {
	RequestId     string
	MediaWorkflow ActivateMediaWorkflowMediaWorkflow
}

type ActivateMediaWorkflowMediaWorkflow struct {
	MediaWorkflowId string
	Name            string
	Topology        string
	State           string
	CreationTime    string
}
