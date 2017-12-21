package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type HeadquartersOverviewRequest struct {
	Bid int64 `position:"Query" name:"Bid"`
}

func (r HeadquartersOverviewRequest) Invoke(client *sdk.Client) (response *HeadquartersOverviewResponse, err error) {
	req := struct {
		*requests.RpcRequest
		HeadquartersOverviewRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "HeadquartersOverview", "", "")

	resp := struct {
		*responses.BaseResponse
		HeadquartersOverviewResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.HeadquartersOverviewResponse

	err = client.DoAction(&req, &resp)
	return
}

type HeadquartersOverviewResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
