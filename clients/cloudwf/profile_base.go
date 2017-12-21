package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ProfileBaseRequest struct {
	requests.RpcRequest
	Gsid int64 `position:"Query" name:"Gsid"`
}

func (req *ProfileBaseRequest) Invoke(client *sdk.Client) (resp *ProfileBaseResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ProfileBase", "", "")
	resp = &ProfileBaseResponse{}
	err = client.DoAction(req, resp)
	return
}

type ProfileBaseResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
