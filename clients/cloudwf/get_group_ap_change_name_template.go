package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetGroupApChangeNameTemplateRequest struct {
	requests.RpcRequest
}

func (req *GetGroupApChangeNameTemplateRequest) Invoke(client *sdk.Client) (resp *GetGroupApChangeNameTemplateResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetGroupApChangeNameTemplate", "", "")
	resp = &GetGroupApChangeNameTemplateResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetGroupApChangeNameTemplateResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
