
package r-kvstore

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)


type RenewInstanceRequest struct {
ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
Period int64 `position:"Query" name:"Period"`
AutoPay string `position:"Query" name:"AutoPay"`
FromApp string `position:"Query" name:"FromApp"`
ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
OwnerAccount string `position:"Query" name:"OwnerAccount"`
CouponNo string `position:"Query" name:"CouponNo"`
OwnerId int64 `position:"Query" name:"OwnerId"`
InstanceClass string `position:"Query" name:"InstanceClass"`
Capacity string `position:"Query" name:"Capacity"`
InstanceId string `position:"Query" name:"InstanceId"`
SecurityToken string `position:"Query" name:"SecurityToken"`
ForceUpgrade string `position:"Query" name:"ForceUpgrade"`
BusinessInfo string `position:"Query" name:"BusinessInfo"`
}

func (r RenewInstanceRequest) Invoke(client *sdk.Client) (response *RenewInstanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RenewInstanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "RenewInstance", "redisa", "")

	resp := struct {
		*responses.BaseResponse
		RenewInstanceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
    response = &resp.RenewInstanceResponse

	err = client.DoAction(&req, &resp)
	return
}

type RenewInstanceResponse struct {
RequestId string
OrderId string
EndTime string
}

