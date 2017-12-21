package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type BusinessInfoRequest struct {
	Bid int64 `position:"Query" name:"Bid"`
}

func (r BusinessInfoRequest) Invoke(client *sdk.Client) (response *BusinessInfoResponse, err error) {
	req := struct {
		*requests.RpcRequest
		BusinessInfoRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "BusinessInfo", "", "")

	resp := struct {
		*responses.BaseResponse
		BusinessInfoResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.BusinessInfoResponse

	err = client.DoAction(&req, &resp)
	return
}

type BusinessInfoResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
