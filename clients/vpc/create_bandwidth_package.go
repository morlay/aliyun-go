package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateBandwidthPackageRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	Bandwidth            int    `position:"Query" name:"Bandwidth"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	ISP                  string `position:"Query" name:"ISP"`
	Description          string `position:"Query" name:"Description"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Zone                 string `position:"Query" name:"Zone"`
	InternetChargeType   string `position:"Query" name:"InternetChargeType"`
	Name                 string `position:"Query" name:"Name"`
	NatGatewayId         string `position:"Query" name:"NatGatewayId"`
	IpCount              int    `position:"Query" name:"IpCount"`
}

func (r CreateBandwidthPackageRequest) Invoke(client *sdk.Client) (response *CreateBandwidthPackageResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateBandwidthPackageRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "CreateBandwidthPackage", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		CreateBandwidthPackageResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateBandwidthPackageResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateBandwidthPackageResponse struct {
	RequestId          string
	BandwidthPackageId string
}
