package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ShopMarketingListRequest struct {
	requests.RpcRequest
	Name string `position:"Query" name:"Name"`
	Page int    `position:"Query" name:"Page"`
	Per  int    `position:"Query" name:"Per"`
	Sid  int64  `position:"Query" name:"Sid"`
}

func (req *ShopMarketingListRequest) Invoke(client *sdk.Client) (resp *ShopMarketingListResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ShopMarketingList", "", "")
	resp = &ShopMarketingListResponse{}
	err = client.DoAction(req, resp)
	return
}

type ShopMarketingListResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
