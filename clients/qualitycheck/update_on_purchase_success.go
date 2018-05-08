package qualitycheck

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type UpdateOnPurchaseSuccessRequest struct {
	requests.RpcRequest
	Data string `position:"Query" name:"Data"`
}

func (req *UpdateOnPurchaseSuccessRequest) Invoke(client *sdk.Client) (resp *UpdateOnPurchaseSuccessResponse, err error) {
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "UpdateOnPurchaseSuccess", "", "")
	resp = &UpdateOnPurchaseSuccessResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateOnPurchaseSuccessResponse struct {
	responses.BaseResponse
	RequestId common.String
	Data      common.String
	Success   bool
	Code      common.String
	Message   common.String
}
