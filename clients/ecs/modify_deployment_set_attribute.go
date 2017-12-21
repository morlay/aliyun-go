package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyDeploymentSetAttributeRequest struct {
	DeploymentSetId      string `position:"Query" name:"DeploymentSetId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Description          string `position:"Query" name:"Description"`
	DeploymentSetName    string `position:"Query" name:"DeploymentSetName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ModifyDeploymentSetAttributeRequest) Invoke(client *sdk.Client) (response *ModifyDeploymentSetAttributeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyDeploymentSetAttributeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "ModifyDeploymentSetAttribute", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		ModifyDeploymentSetAttributeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyDeploymentSetAttributeResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyDeploymentSetAttributeResponse struct {
	RequestId string
}
