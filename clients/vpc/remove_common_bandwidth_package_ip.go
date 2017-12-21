package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RemoveCommonBandwidthPackageIpRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	BandwidthPackageId   string `position:"Query" name:"BandwidthPackageId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	IpInstanceId         string `position:"Query" name:"IpInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r RemoveCommonBandwidthPackageIpRequest) Invoke(client *sdk.Client) (response *RemoveCommonBandwidthPackageIpResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RemoveCommonBandwidthPackageIpRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "RemoveCommonBandwidthPackageIp", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		RemoveCommonBandwidthPackageIpResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RemoveCommonBandwidthPackageIpResponse

	err = client.DoAction(&req, &resp)
	return
}

type RemoveCommonBandwidthPackageIpResponse struct {
	RequestId string
}
