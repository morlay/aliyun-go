package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type FindApRequest struct {
	requests.RpcRequest
	Id int64 `position:"Query" name:"Id"`
}

func (req *FindApRequest) Invoke(client *sdk.Client) (resp *FindApResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "FindAp", "", "")
	resp = &FindApResponse{}
	err = client.DoAction(req, resp)
	return
}

type FindApResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
