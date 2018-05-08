package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteDeploymentSetRequest struct {
	requests.RpcRequest
	DeploymentSetId      string `position:"Query" name:"DeploymentSetId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteDeploymentSetRequest) Invoke(client *sdk.Client) (resp *DeleteDeploymentSetResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DeleteDeploymentSet", "ecs", "")
	resp = &DeleteDeploymentSetResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteDeploymentSetResponse struct {
	responses.BaseResponse
	RequestId common.String
}
