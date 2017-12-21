package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteMediaWorkflowRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	MediaWorkflowId      string `position:"Query" name:"MediaWorkflowId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteMediaWorkflowRequest) Invoke(client *sdk.Client) (resp *DeleteMediaWorkflowResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "DeleteMediaWorkflow", "", "")
	resp = &DeleteMediaWorkflowResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteMediaWorkflowResponse struct {
	responses.BaseResponse
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
