package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetSubAccountStatusRequest struct {
}

func (r GetSubAccountStatusRequest) Invoke(client *sdk.Client) (response *GetSubAccountStatusResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetSubAccountStatusRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetSubAccountStatus", "", "")

	resp := struct {
		*responses.BaseResponse
		GetSubAccountStatusResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetSubAccountStatusResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetSubAccountStatusResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
