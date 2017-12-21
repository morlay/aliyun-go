package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetGroupApRadioConfigTemplateRequest struct {
}

func (r GetGroupApRadioConfigTemplateRequest) Invoke(client *sdk.Client) (response *GetGroupApRadioConfigTemplateResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetGroupApRadioConfigTemplateRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetGroupApRadioConfigTemplate", "", "")

	resp := struct {
		*responses.BaseResponse
		GetGroupApRadioConfigTemplateResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetGroupApRadioConfigTemplateResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetGroupApRadioConfigTemplateResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
