package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyCommonBandwidthPackageSpecRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	BandwidthPackageId   string `position:"Query" name:"BandwidthPackageId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth            string `position:"Query" name:"Bandwidth"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifyCommonBandwidthPackageSpecRequest) Invoke(client *sdk.Client) (resp *ModifyCommonBandwidthPackageSpecResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "ModifyCommonBandwidthPackageSpec", "vpc", "")
	resp = &ModifyCommonBandwidthPackageSpecResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyCommonBandwidthPackageSpecResponse struct {
	responses.BaseResponse
	RequestId string
}
