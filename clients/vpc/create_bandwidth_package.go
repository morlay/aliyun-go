package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateBandwidthPackageRequest struct {
	requests.RpcRequest
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

func (req *CreateBandwidthPackageRequest) Invoke(client *sdk.Client) (resp *CreateBandwidthPackageResponse, err error) {
	req.InitWithApiInfo("Vpc", "2016-04-28", "CreateBandwidthPackage", "vpc", "")
	resp = &CreateBandwidthPackageResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateBandwidthPackageResponse struct {
	responses.BaseResponse
	RequestId          common.String
	BandwidthPackageId common.String
}
