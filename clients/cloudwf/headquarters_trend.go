package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type HeadquartersTrendRequest struct {
	Bid int64 `position:"Query" name:"Bid"`
}

func (r HeadquartersTrendRequest) Invoke(client *sdk.Client) (response *HeadquartersTrendResponse, err error) {
	req := struct {
		*requests.RpcRequest
		HeadquartersTrendRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "HeadquartersTrend", "", "")

	resp := struct {
		*responses.BaseResponse
		HeadquartersTrendResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.HeadquartersTrendResponse

	err = client.DoAction(&req, &resp)
	return
}

type HeadquartersTrendResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
