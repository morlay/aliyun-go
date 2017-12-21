package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GroupOverviewRequest struct {
	Gsid int64 `position:"Query" name:"Gsid"`
}

func (r GroupOverviewRequest) Invoke(client *sdk.Client) (response *GroupOverviewResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GroupOverviewRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "GroupOverview", "", "")

	resp := struct {
		*responses.BaseResponse
		GroupOverviewResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GroupOverviewResponse

	err = client.DoAction(&req, &resp)
	return
}

type GroupOverviewResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
