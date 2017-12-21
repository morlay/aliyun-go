package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateCommonBandwidthPackageRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	Bandwidth            int    `position:"Query" name:"Bandwidth"`
	InternetChargeType   string `position:"Query" name:"InternetChargeType"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	Name                 string `position:"Query" name:"Name"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Ratio                int    `position:"Query" name:"Ratio"`
}

func (r CreateCommonBandwidthPackageRequest) Invoke(client *sdk.Client) (response *CreateCommonBandwidthPackageResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateCommonBandwidthPackageRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "CreateCommonBandwidthPackage", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		CreateCommonBandwidthPackageResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateCommonBandwidthPackageResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateCommonBandwidthPackageResponse struct {
	RequestId          string
	BandwidthPackageId string
}
