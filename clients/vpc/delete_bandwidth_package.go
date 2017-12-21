package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteBandwidthPackageRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	BandwidthPackageId   string `position:"Query" name:"BandwidthPackageId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Force                string `position:"Query" name:"Force"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DeleteBandwidthPackageRequest) Invoke(client *sdk.Client) (response *DeleteBandwidthPackageResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteBandwidthPackageRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "DeleteBandwidthPackage", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		DeleteBandwidthPackageResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteBandwidthPackageResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteBandwidthPackageResponse struct {
	RequestId string
}
