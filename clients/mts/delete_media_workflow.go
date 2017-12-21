package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteMediaWorkflowRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	MediaWorkflowId      string `position:"Query" name:"MediaWorkflowId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DeleteMediaWorkflowRequest) Invoke(client *sdk.Client) (response *DeleteMediaWorkflowResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteMediaWorkflowRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Mts", "2014-06-18", "DeleteMediaWorkflow", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteMediaWorkflowResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteMediaWorkflowResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteMediaWorkflowResponse struct {
	RequestId     string
	MediaWorkflow DeleteMediaWorkflowMediaWorkflow
}

type DeleteMediaWorkflowMediaWorkflow struct {
	MediaWorkflowId string
	Name            string
	Topology        string
	State           string
	CreationTime    string
}
