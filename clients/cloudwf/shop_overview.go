package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ShopOverviewRequest struct {
	Gsid int64 `position:"Query" name:"Gsid"`
}

func (r ShopOverviewRequest) Invoke(client *sdk.Client) (response *ShopOverviewResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ShopOverviewRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ShopOverview", "", "")

	resp := struct {
		*responses.BaseResponse
		ShopOverviewResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ShopOverviewResponse

	err = client.DoAction(&req, &resp)
	return
}

type ShopOverviewResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
