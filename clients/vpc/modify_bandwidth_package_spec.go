package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyBandwidthPackageSpecRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	BandwidthPackageId   string `position:"Query" name:"BandwidthPackageId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth            string `position:"Query" name:"Bandwidth"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r ModifyBandwidthPackageSpecRequest) Invoke(client *sdk.Client) (response *ModifyBandwidthPackageSpecResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyBandwidthPackageSpecRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "ModifyBandwidthPackageSpec", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		ModifyBandwidthPackageSpecResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyBandwidthPackageSpecResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyBandwidthPackageSpecResponse struct {
	RequestId string
}
