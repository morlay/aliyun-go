package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ResetApRequest struct {
	requests.RpcRequest
	Id int64 `position:"Query" name:"Id"`
}

func (req *ResetApRequest) Invoke(client *sdk.Client) (resp *ResetApResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ResetAp", "", "")
	resp = &ResetApResponse{}
	err = client.DoAction(req, resp)
	return
}

type ResetApResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
