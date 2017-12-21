
package r-kvstore

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)


type RenewMultiInstanceRequest struct {
ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
Period int64 `position:"Query" name:"Period"`
AutoPay string `position:"Query" name:"AutoPay"`
FromApp string `position:"Query" name:"FromApp"`
ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
OwnerAccount string `position:"Query" name:"OwnerAccount"`
CouponNo string `position:"Query" name:"CouponNo"`
OwnerId int64 `position:"Query" name:"OwnerId"`
SecurityToken string `position:"Query" name:"SecurityToken"`
InstanceIds string `position:"Query" name:"InstanceIds"`
BusinessInfo string `position:"Query" name:"BusinessInfo"`
}

func (r RenewMultiInstanceRequest) Invoke(client *sdk.Client) (response *RenewMultiInstanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RenewMultiInstanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "RenewMultiInstance", "redisa", "")

	resp := struct {
		*responses.BaseResponse
		RenewMultiInstanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
    response = &resp.RenewMultiInstanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type RenewMultiInstanceResponse struct {
RequestId string
OrderId string
}

