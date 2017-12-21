
package r-kvstore

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
)


type TransformToPrePaidRequest struct {
ResourceOwnerId int64 `position:"Query" name:"ResourceOwnerId"`
Period int64 `position:"Query" name:"Period"`
InstanceId string `position:"Query" name:"InstanceId"`
AutoPay string `position:"Query" name:"AutoPay"`
FromApp string `position:"Query" name:"FromApp"`
SecurityToken string `position:"Query" name:"SecurityToken"`
ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
OwnerAccount string `position:"Query" name:"OwnerAccount"`
OwnerId int64 `position:"Query" name:"OwnerId"`
}

func (r TransformToPrePaidRequest) Invoke(client *sdk.Client) (response *TransformToPrePaidResponse, err error) {
	req := struct {
		*requests.RpcRequest
		TransformToPrePaidRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("R-kvstore", "2015-01-01", "TransformToPrePaid", "redisa", "")

	resp := struct {
		*responses.BaseResponse
		TransformToPrePaidResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
    response = &resp.TransformToPrePaidResponse

	err = client.DoAction(&req, &resp)
	return
}

type TransformToPrePaidResponse struct {
RequestId string
OrderId string
EndTime string
}

