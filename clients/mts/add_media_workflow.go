package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AddMediaWorkflowRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Topology             string `position:"Query" name:"Topology"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	TriggerMode          string `position:"Query" name:"TriggerMode"`
}

func (req *AddMediaWorkflowRequest) Invoke(client *sdk.Client) (resp *AddMediaWorkflowResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "AddMediaWorkflow", "mts", "")
	resp = &AddMediaWorkflowResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddMediaWorkflowResponse struct {
	responses.BaseResponse
	RequestId     common.String
	MediaWorkflow AddMediaWorkflowMediaWorkflow
}

type AddMediaWorkflowMediaWorkflow struct {
	MediaWorkflowId common.String
	Name            common.String
	Topology        common.String
	TriggerMode     common.String
	State           common.String
	CreationTime    common.String
}
