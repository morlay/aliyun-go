package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetAccountConfigRequest struct {
	requests.RpcRequest
	Id int64 `position:"Query" name:"Id"`
}

func (req *GetAccountConfigRequest) Invoke(client *sdk.Client) (resp *GetAccountConfigResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetAccountConfig", "", "")
	resp = &GetAccountConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetAccountConfigResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
