package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetSubAccountStatusRequest struct {
	requests.RpcRequest
}

func (req *GetSubAccountStatusRequest) Invoke(client *sdk.Client) (resp *GetSubAccountStatusResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetSubAccountStatus", "", "")
	resp = &GetSubAccountStatusResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetSubAccountStatusResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
