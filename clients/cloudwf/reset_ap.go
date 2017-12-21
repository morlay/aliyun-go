package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ResetApRequest struct {
	Id int64 `position:"Query" name:"Id"`
}

func (r ResetApRequest) Invoke(client *sdk.Client) (response *ResetApResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ResetApRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ResetAp", "", "")

	resp := struct {
		*responses.BaseResponse
		ResetApResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ResetApResponse

	err = client.DoAction(&req, &resp)
	return
}

type ResetApResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
