package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteHpcClusterRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	HpcClusterId         string `position:"Query" name:"HpcClusterId"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DeleteHpcClusterRequest) Invoke(client *sdk.Client) (response *DeleteHpcClusterResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteHpcClusterRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "DeleteHpcCluster", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		DeleteHpcClusterResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteHpcClusterResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteHpcClusterResponse struct {
	RequestId string
}
