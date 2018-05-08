package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteBandwidthPackageRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	BandwidthPackageId   string `position:"Query" name:"BandwidthPackageId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteBandwidthPackageRequest) Invoke(client *sdk.Client) (resp *DeleteBandwidthPackageResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DeleteBandwidthPackage", "ecs", "")
	resp = &DeleteBandwidthPackageResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteBandwidthPackageResponse struct {
	responses.BaseResponse
	RequestId common.String
}
