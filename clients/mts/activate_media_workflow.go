package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ActivateMediaWorkflowRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	MediaWorkflowId      string `position:"Query" name:"MediaWorkflowId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ActivateMediaWorkflowRequest) Invoke(client *sdk.Client) (resp *ActivateMediaWorkflowResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "ActivateMediaWorkflow", "mts", "")
	resp = &ActivateMediaWorkflowResponse{}
	err = client.DoAction(req, resp)
	return
}

type ActivateMediaWorkflowResponse struct {
	responses.BaseResponse
	RequestId     common.String
	MediaWorkflow ActivateMediaWorkflowMediaWorkflow
}

type ActivateMediaWorkflowMediaWorkflow struct {
	MediaWorkflowId common.String
	Name            common.String
	Topology        common.String
	State           common.String
	CreationTime    common.String
}
