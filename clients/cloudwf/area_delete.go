package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AreaDeleteRequest struct {
	Aid int64 `position:"Query" name:"Aid"`
	Sid int64 `position:"Query" name:"Sid"`
}

func (r AreaDeleteRequest) Invoke(client *sdk.Client) (response *AreaDeleteResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AreaDeleteRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "AreaDelete", "", "")

	resp := struct {
		*responses.BaseResponse
		AreaDeleteResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AreaDeleteResponse

	err = client.DoAction(&req, &resp)
	return
}

type AreaDeleteResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
