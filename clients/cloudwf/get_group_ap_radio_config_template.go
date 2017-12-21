package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetGroupApRadioConfigTemplateRequest struct {
	requests.RpcRequest
}

func (req *GetGroupApRadioConfigTemplateRequest) Invoke(client *sdk.Client) (resp *GetGroupApRadioConfigTemplateResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetGroupApRadioConfigTemplate", "", "")
	resp = &GetGroupApRadioConfigTemplateResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetGroupApRadioConfigTemplateResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
