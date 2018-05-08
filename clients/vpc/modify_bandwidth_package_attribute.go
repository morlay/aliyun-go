package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyBandwidthPackageAttributeRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	BandwidthPackageId   string `position:"Query" name:"BandwidthPackageId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifyBandwidthPackageAttributeRequest) Invoke(client *sdk.Client) (resp *ModifyBandwidthPackageAttributeResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "ModifyBandwidthPackageAttribute", "vpc", "")
	resp = &ModifyBandwidthPackageAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyBandwidthPackageAttributeResponse struct {
	responses.BaseResponse
	RequestId common.String
}
