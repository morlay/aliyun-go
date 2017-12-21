package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetAccountConfigRequest struct {
	Id int64 `position:"Query" name:"Id"`
}

func (r GetAccountConfigRequest) Invoke(client *sdk.Client) (response *GetAccountConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetAccountConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetAccountConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		GetAccountConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetAccountConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetAccountConfigResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
