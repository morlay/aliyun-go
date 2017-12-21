package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteBandwidthPackageRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	BandwidthPackageId   string `position:"Query" name:"BandwidthPackageId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Force                string `position:"Query" name:"Force"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteBandwidthPackageRequest) Invoke(client *sdk.Client) (resp *DeleteBandwidthPackageResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "DeleteBandwidthPackage", "vpc", "")
	resp = &DeleteBandwidthPackageResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteBandwidthPackageResponse struct {
	responses.BaseResponse
	RequestId string
}
