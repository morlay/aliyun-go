package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type UpdateMediaWorkflowTriggerModeRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	MediaWorkflowId      string `position:"Query" name:"MediaWorkflowId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	TriggerMode          string `position:"Query" name:"TriggerMode"`
}

func (req *UpdateMediaWorkflowTriggerModeRequest) Invoke(client *sdk.Client) (resp *UpdateMediaWorkflowTriggerModeResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "UpdateMediaWorkflowTriggerMode", "mts", "")
	resp = &UpdateMediaWorkflowTriggerModeResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateMediaWorkflowTriggerModeResponse struct {
	responses.BaseResponse
	RequestId     common.String
	MediaWorkflow UpdateMediaWorkflowTriggerModeMediaWorkflow
}

type UpdateMediaWorkflowTriggerModeMediaWorkflow struct {
	MediaWorkflowId common.String
	Name            common.String
	Topology        common.String
	TriggerMode     common.String
	State           common.String
	CreationTime    common.String
}
