package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetBindAp4UmengRequest struct {
	requests.RpcRequest
}

func (req *GetBindAp4UmengRequest) Invoke(client *sdk.Client) (resp *GetBindAp4UmengResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetBindAp4Umeng", "", "")
	resp = &GetBindAp4UmengResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetBindAp4UmengResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
