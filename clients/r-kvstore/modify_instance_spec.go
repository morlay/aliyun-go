
package r-kvstore

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)


type ModifyInstanceSpecRequest struct {
ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
FromApp string `position:"Query" name:"FromApp"`
ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
OwnerAccount string `position:"Query" name:"OwnerAccount"`
CouponNo string `position:"Query" name:"CouponNo"`
OwnerId int64 `position:"Query" name:"OwnerId"`
InstanceClass string `position:"Query" name:"InstanceClass"`
InstanceId string `position:"Query" name:"InstanceId"`
SecurityToken string `position:"Query" name:"SecurityToken"`
ForceUpgrade string `position:"Query" name:"ForceUpgrade"`
BusinessInfo string `position:"Query" name:"BusinessInfo"`
}

func (r ModifyInstanceSpecRequest) Invoke(client *sdk.Client) (response *ModifyInstanceSpecResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyInstanceSpecRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "ModifyInstanceSpec", "redisa", "")

	resp := struct {
		*responses.BaseResponse
		ModifyInstanceSpecResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
    response = &resp.ModifyInstanceSpecResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyInstanceSpecResponse struct {
RequestId string
OrderId string
}

