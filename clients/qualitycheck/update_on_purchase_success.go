package qualitycheck

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateOnPurchaseSuccessRequest struct {
	Data string `position:"Query" name:"Data"`
}

func (r UpdateOnPurchaseSuccessRequest) Invoke(client *sdk.Client) (response *UpdateOnPurchaseSuccessResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UpdateOnPurchaseSuccessRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "UpdateOnPurchaseSuccess", "", "")

	resp := struct {
		*responses.BaseResponse
		UpdateOnPurchaseSuccessResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UpdateOnPurchaseSuccessResponse

	err = client.DoAction(&req, &resp)
	return
}

type UpdateOnPurchaseSuccessResponse struct {
	RequestId string
	Data      string
	Success   bool
	Code      string
	Message   string
}
