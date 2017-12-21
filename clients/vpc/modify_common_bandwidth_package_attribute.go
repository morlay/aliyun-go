package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyCommonBandwidthPackageAttributeRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	BandwidthPackageId   string `position:"Query" name:"BandwidthPackageId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifyCommonBandwidthPackageAttributeRequest) Invoke(client *sdk.Client) (resp *ModifyCommonBandwidthPackageAttributeResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "ModifyCommonBandwidthPackageAttribute", "vpc", "")
	resp = &ModifyCommonBandwidthPackageAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyCommonBandwidthPackageAttributeResponse struct {
	responses.BaseResponse
	RequestId string
}
