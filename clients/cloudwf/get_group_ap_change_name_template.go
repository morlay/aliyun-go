package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetGroupApChangeNameTemplateRequest struct {
}

func (r GetGroupApChangeNameTemplateRequest) Invoke(client *sdk.Client) (response *GetGroupApChangeNameTemplateResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetGroupApChangeNameTemplateRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetGroupApChangeNameTemplate", "", "")

	resp := struct {
		*responses.BaseResponse
		GetGroupApChangeNameTemplateResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetGroupApChangeNameTemplateResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetGroupApChangeNameTemplateResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
