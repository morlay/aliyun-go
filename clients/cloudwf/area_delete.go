package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AreaDeleteRequest struct {
	requests.RpcRequest
	Aid int64 `position:"Query" name:"Aid"`
	Sid int64 `position:"Query" name:"Sid"`
}

func (req *AreaDeleteRequest) Invoke(client *sdk.Client) (resp *AreaDeleteResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "AreaDelete", "", "")
	resp = &AreaDeleteResponse{}
	err = client.DoAction(req, resp)
	return
}

type AreaDeleteResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
