package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ShopCameraRequest struct {
	Gsid int64 `position:"Query" name:"Gsid"`
}

func (r ShopCameraRequest) Invoke(client *sdk.Client) (response *ShopCameraResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ShopCameraRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ShopCamera", "", "")

	resp := struct {
		*responses.BaseResponse
		ShopCameraResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ShopCameraResponse

	err = client.DoAction(&req, &resp)
	return
}

type ShopCameraResponse struct {
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
