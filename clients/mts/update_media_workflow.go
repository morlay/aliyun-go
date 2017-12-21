package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateMediaWorkflowRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Topology             string `position:"Query" name:"Topology"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	MediaWorkflowId      string `position:"Query" name:"MediaWorkflowId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *UpdateMediaWorkflowRequest) Invoke(client *sdk.Client) (resp *UpdateMediaWorkflowResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "UpdateMediaWorkflow", "", "")
	resp = &UpdateMediaWorkflowResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateMediaWorkflowResponse struct {
	responses.BaseResponse
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
