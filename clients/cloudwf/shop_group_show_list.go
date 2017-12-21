package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ShopGroupShowListRequest struct {
	requests.RpcRequest
	Page int   `position:"Query" name:"Page"`
	Bid  int64 `position:"Query" name:"Bid"`
	Per  int   `position:"Query" name:"Per"`
}

func (req *ShopGroupShowListRequest) Invoke(client *sdk.Client) (resp *ShopGroupShowListResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ShopGroupShowList", "", "")
	resp = &ShopGroupShowListResponse{}
	err = client.DoAction(req, resp)
	return
}

type ShopGroupShowListResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
