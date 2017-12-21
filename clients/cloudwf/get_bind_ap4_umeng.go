package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetBindAp4UmengRequest struct {
}

func (r GetBindAp4UmengRequest) Invoke(client *sdk.Client) (response *GetBindAp4UmengResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetBindAp4UmengRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetBindAp4Umeng", "", "")

	resp := struct {
		*responses.BaseResponse
		GetBindAp4UmengResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetBindAp4UmengResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetBindAp4UmengResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
