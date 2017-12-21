package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ProfileConsumeRequest struct {
	requests.RpcRequest
	Gsid int64 `position:"Query" name:"Gsid"`
}

func (req *ProfileConsumeRequest) Invoke(client *sdk.Client) (resp *ProfileConsumeResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ProfileConsume", "", "")
	resp = &ProfileConsumeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ProfileConsumeResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
