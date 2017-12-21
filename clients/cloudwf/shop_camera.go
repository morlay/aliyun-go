package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ShopCameraRequest struct {
	requests.RpcRequest
	Gsid int64 `position:"Query" name:"Gsid"`
}

func (req *ShopCameraRequest) Invoke(client *sdk.Client) (resp *ShopCameraResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ShopCamera", "", "")
	resp = &ShopCameraResponse{}
	err = client.DoAction(req, resp)
	return
}

type ShopCameraResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
