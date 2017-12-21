package mts

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddMediaWorkflowRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Topology             string `position:"Query" name:"Topology"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *AddMediaWorkflowRequest) Invoke(client *sdk.Client) (resp *AddMediaWorkflowResponse, err error) {
	req.InitWithApiInfo("Mts", "2014-06-18", "AddMediaWorkflow", "", "")
	resp = &AddMediaWorkflowResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddMediaWorkflowResponse struct {
	responses.BaseResponse
	RequestId     string
	MediaWorkflow AddMediaWorkflowMediaWorkflow
}

type AddMediaWorkflowMediaWorkflow struct {
	MediaWorkflowId string
	Name            string
	Topology        string
	State           string
	CreationTime    string
}
