package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeactivateMediaWorkflowRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	MediaWorkflowId      string `position:"Query" name:"MediaWorkflowId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DeactivateMediaWorkflowRequest) Invoke(client *sdk.Client) (resp *DeactivateMediaWorkflowResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "DeactivateMediaWorkflow", "mts", "")
	resp = &DeactivateMediaWorkflowResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeactivateMediaWorkflowResponse struct {
	responses.BaseResponse
	RequestId     common.String
	MediaWorkflow DeactivateMediaWorkflowMediaWorkflow
}

type DeactivateMediaWorkflowMediaWorkflow struct {
	MediaWorkflowId common.String
	Name            common.String
	Topology        common.String
	State           common.String
	CreationTime    common.String
}
