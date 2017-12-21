package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AreaInfoRequest struct {
	requests.RpcRequest
	Aid int64 `position:"Query" name:"Aid"`
	Sid int64 `position:"Query" name:"Sid"`
}

func (req *AreaInfoRequest) Invoke(client *sdk.Client) (resp *AreaInfoResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "AreaInfo", "", "")
	resp = &AreaInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type AreaInfoResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
