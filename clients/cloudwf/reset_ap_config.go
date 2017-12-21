package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ResetApConfigRequest struct {
	Id int64 `position:"Query" name:"Id"`
}

func (r ResetApConfigRequest) Invoke(client *sdk.Client) (response *ResetApConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ResetApConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ResetApConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		ResetApConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ResetApConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type ResetApConfigResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
