package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateDeploymentSetRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	DeploymentSetName    string `position:"Query" name:"DeploymentSetName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Granularity          string `position:"Query" name:"Granularity"`
	Domain               string `position:"Query" name:"Domain"`
	ZoneId               string `position:"Query" name:"ZoneId"`
	Strategy             string `position:"Query" name:"Strategy"`
}

func (req *CreateDeploymentSetRequest) Invoke(client *sdk.Client) (resp *CreateDeploymentSetResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "CreateDeploymentSet", "ecs", "")
	resp = &CreateDeploymentSetResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateDeploymentSetResponse struct {
	responses.BaseResponse
	RequestId       string
	DeploymentSetId string
}
