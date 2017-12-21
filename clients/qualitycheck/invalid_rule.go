package qualitycheck

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type InvalidRuleRequest struct {
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (r InvalidRuleRequest) Invoke(client *sdk.Client) (response *InvalidRuleResponse, err error) {
	req := struct {
		*requests.RpcRequest
		InvalidRuleRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "InvalidRule", "", "")

	resp := struct {
		*responses.BaseResponse
		InvalidRuleResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.InvalidRuleResponse

	err = client.DoAction(&req, &resp)
	return
}

type InvalidRuleResponse struct {
	RequestId string
	Success   bool
	Code      string
	Message   string
	Data      bool
}
