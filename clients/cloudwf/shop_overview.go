package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ShopOverviewRequest struct {
	requests.RpcRequest
	Gsid int64 `position:"Query" name:"Gsid"`
}

func (req *ShopOverviewRequest) Invoke(client *sdk.Client) (resp *ShopOverviewResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ShopOverview", "", "")
	resp = &ShopOverviewResponse{}
	err = client.DoAction(req, resp)
	return
}

type ShopOverviewResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
