package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ShopMarketingListRequest struct {
	Name string `position:"Query" name:"Name"`
	Page int    `position:"Query" name:"Page"`
	Per  int    `position:"Query" name:"Per"`
	Sid  int64  `position:"Query" name:"Sid"`
}

func (r ShopMarketingListRequest) Invoke(client *sdk.Client) (response *ShopMarketingListResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ShopMarketingListRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ShopMarketingList", "", "")

	resp := struct {
		*responses.BaseResponse
		ShopMarketingListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ShopMarketingListResponse

	err = client.DoAction(&req, &resp)
	return
}

type ShopMarketingListResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
