package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetBidRequest struct {
}

func (r GetBidRequest) Invoke(client *sdk.Client) (response *GetBidResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetBidRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetBid", "", "")

	resp := struct {
		*responses.BaseResponse
		GetBidResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetBidResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetBidResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
