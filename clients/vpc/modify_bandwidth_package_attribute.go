package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyBandwidthPackageAttributeRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	BandwidthPackageId   string `position:"Query" name:"BandwidthPackageId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ModifyBandwidthPackageAttributeRequest) Invoke(client *sdk.Client) (response *ModifyBandwidthPackageAttributeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyBandwidthPackageAttributeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "ModifyBandwidthPackageAttribute", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		ModifyBandwidthPackageAttributeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyBandwidthPackageAttributeResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyBandwidthPackageAttributeResponse struct {
	RequestId string
}
