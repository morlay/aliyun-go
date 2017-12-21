package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ProfileTradeRequest struct {
	Gsid int64 `position:"Query" name:"Gsid"`
}

func (r ProfileTradeRequest) Invoke(client *sdk.Client) (response *ProfileTradeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ProfileTradeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ProfileTrade", "", "")

	resp := struct {
		*responses.BaseResponse
		ProfileTradeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ProfileTradeResponse

	err = client.DoAction(&req, &resp)
	return
}

type ProfileTradeResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
