package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyCommonBandwidthPackageAttributeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	BandwidthPackageId   string `position:"Query" name:"BandwidthPackageId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ModifyCommonBandwidthPackageAttributeRequest) Invoke(client *sdk.Client) (response *ModifyCommonBandwidthPackageAttributeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyCommonBandwidthPackageAttributeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "ModifyCommonBandwidthPackageAttribute", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		ModifyCommonBandwidthPackageAttributeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyCommonBandwidthPackageAttributeResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyCommonBandwidthPackageAttributeResponse struct {
	RequestId string
}
