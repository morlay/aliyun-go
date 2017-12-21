package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ShopDeleteRequest struct {
	requests.RpcRequest
	Sid int64 `position:"Query" name:"Sid"`
}

func (req *ShopDeleteRequest) Invoke(client *sdk.Client) (resp *ShopDeleteResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ShopDelete", "", "")
	resp = &ShopDeleteResponse{}
	err = client.DoAction(req, resp)
	return
}

type ShopDeleteResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
