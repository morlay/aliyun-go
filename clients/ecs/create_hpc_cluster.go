package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateHpcClusterRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	Description          string `position:"Query" name:"Description"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Name                 string `position:"Query" name:"Name"`
}

func (r CreateHpcClusterRequest) Invoke(client *sdk.Client) (response *CreateHpcClusterResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateHpcClusterRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "CreateHpcCluster", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		CreateHpcClusterResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateHpcClusterResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateHpcClusterResponse struct {
	RequestId    string
	HpcClusterId string
}
