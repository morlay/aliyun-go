package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RemoveCommonBandwidthPackageIpRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	BandwidthPackageId   string `position:"Query" name:"BandwidthPackageId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	IpInstanceId         string `position:"Query" name:"IpInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *RemoveCommonBandwidthPackageIpRequest) Invoke(client *sdk.Client) (resp *RemoveCommonBandwidthPackageIpResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "RemoveCommonBandwidthPackageIp", "vpc", "")
	resp = &RemoveCommonBandwidthPackageIpResponse{}
	err = client.DoAction(req, resp)
	return
}

type RemoveCommonBandwidthPackageIpResponse struct {
	responses.BaseResponse
	RequestId string
}
