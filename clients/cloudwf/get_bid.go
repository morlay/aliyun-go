package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetBidRequest struct {
	requests.RpcRequest
}

func (req *GetBidRequest) Invoke(client *sdk.Client) (resp *GetBidResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GetBid", "", "")
	resp = &GetBidResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetBidResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
