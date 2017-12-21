package aegis

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RenewInstanceRequest struct {
	Duration     int    `position:"Query" name:"Duration"`
	InstanceId   string `position:"Query" name:"InstanceId"`
	ClientToken  string `position:"Query" name:"ClientToken"`
	OwnerId      int64  `position:"Query" name:"OwnerId"`
	PricingCycle string `position:"Query" name:"PricingCycle"`
}

func (r RenewInstanceRequest) Invoke(client *sdk.Client) (response *RenewInstanceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RenewInstanceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("aegis", "2016-11-11", "RenewInstance", "", "")

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
	OrderId   string
	RequestId string
}
