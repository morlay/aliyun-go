package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	req.InitWithApiInfo("Mts", "2014-06-18", "UpdateMediaWorkflow", "mts", "")
	resp = &UpdateMediaWorkflowResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateMediaWorkflowResponse struct {
	responses.BaseResponse
	RequestId     common.String
	MediaWorkflow UpdateMediaWorkflowMediaWorkflow
}

type UpdateMediaWorkflowMediaWorkflow struct {
	MediaWorkflowId common.String
	Name            common.String
	Topology        common.String
	TriggerMode     common.String
	State           common.String
	CreationTime    common.String
}
