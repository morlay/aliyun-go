package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AreaInfoRequest struct {
	Aid int64 `position:"Query" name:"Aid"`
	Sid int64 `position:"Query" name:"Sid"`
}

func (r AreaInfoRequest) Invoke(client *sdk.Client) (response *AreaInfoResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AreaInfoRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "AreaInfo", "", "")

	resp := struct {
		*responses.BaseResponse
		AreaInfoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AreaInfoResponse

	err = client.DoAction(&req, &resp)
	return
}

type AreaInfoResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
