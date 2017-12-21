package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ShopGroupDeleteRequest struct {
	requests.RpcRequest
	Gid int64 `position:"Query" name:"Gid"`
}

func (req *ShopGroupDeleteRequest) Invoke(client *sdk.Client) (resp *ShopGroupDeleteResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ShopGroupDelete", "", "")
	resp = &ShopGroupDeleteResponse{}
	err = client.DoAction(req, resp)
	return
}

type ShopGroupDeleteResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
