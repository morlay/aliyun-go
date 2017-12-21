package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddCommonBandwidthPackageIpRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	BandwidthPackageId   string `position:"Query" name:"BandwidthPackageId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	IpInstanceId         string `position:"Query" name:"IpInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r AddCommonBandwidthPackageIpRequest) Invoke(client *sdk.Client) (response *AddCommonBandwidthPackageIpResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddCommonBandwidthPackageIpRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "AddCommonBandwidthPackageIp", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		AddCommonBandwidthPackageIpResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddCommonBandwidthPackageIpResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddCommonBandwidthPackageIpResponse struct {
	RequestId string
}
