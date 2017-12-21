package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ResetApConfigRequest struct {
	requests.RpcRequest
	Id int64 `position:"Query" name:"Id"`
}

func (req *ResetApConfigRequest) Invoke(client *sdk.Client) (resp *ResetApConfigResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ResetApConfig", "", "")
	resp = &ResetApConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type ResetApConfigResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
