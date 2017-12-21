package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateDeploymentSetRequest struct {
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

func (r CreateDeploymentSetRequest) Invoke(client *sdk.Client) (response *CreateDeploymentSetResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateDeploymentSetRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "CreateDeploymentSet", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		CreateDeploymentSetResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateDeploymentSetResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateDeploymentSetResponse struct {
	RequestId       string
	DeploymentSetId string
}
