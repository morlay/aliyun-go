package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteHpcClusterRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	HpcClusterId         string `position:"Query" name:"HpcClusterId"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteHpcClusterRequest) Invoke(client *sdk.Client) (resp *DeleteHpcClusterResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DeleteHpcCluster", "ecs", "")
	resp = &DeleteHpcClusterResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteHpcClusterResponse struct {
	responses.BaseResponse
	RequestId string
}
