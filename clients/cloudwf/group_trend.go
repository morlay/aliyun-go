package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GroupTrendRequest struct {
	Gsid int64 `position:"Query" name:"Gsid"`
}

func (r GroupTrendRequest) Invoke(client *sdk.Client) (response *GroupTrendResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GroupTrendRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GroupTrend", "", "")

	resp := struct {
		*responses.BaseResponse
		GroupTrendResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GroupTrendResponse

	err = client.DoAction(&req, &resp)
	return
}

type GroupTrendResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
