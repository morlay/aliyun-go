package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyDeploymentSetAttributeRequest struct {
	requests.RpcRequest
	DeploymentSetId      string `position:"Query" name:"DeploymentSetId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	DeploymentSetName    string `position:"Query" name:"DeploymentSetName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifyDeploymentSetAttributeRequest) Invoke(client *sdk.Client) (resp *ModifyDeploymentSetAttributeResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifyDeploymentSetAttribute", "ecs", "")
	resp = &ModifyDeploymentSetAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyDeploymentSetAttributeResponse struct {
	responses.BaseResponse
	RequestId string
}
