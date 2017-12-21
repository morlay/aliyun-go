package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddCommonBandwidthPackageIpRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	BandwidthPackageId   string `position:"Query" name:"BandwidthPackageId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	IpInstanceId         string `position:"Query" name:"IpInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *AddCommonBandwidthPackageIpRequest) Invoke(client *sdk.Client) (resp *AddCommonBandwidthPackageIpResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "AddCommonBandwidthPackageIp", "vpc", "")
	resp = &AddCommonBandwidthPackageIpResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddCommonBandwidthPackageIpResponse struct {
	responses.BaseResponse
	RequestId string
}
