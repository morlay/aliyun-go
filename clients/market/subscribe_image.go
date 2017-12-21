package market

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SubscribeImageRequest struct {
	ProductCode string `position:"Query" name:"ProductCode"`
}

func (r SubscribeImageRequest) Invoke(client *sdk.Client) (response *SubscribeImageResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SubscribeImageRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Market", "2015-11-01", "SubscribeImage", "", "")

	resp := struct {
		*responses.BaseResponse
		SubscribeImageResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SubscribeImageResponse

	err = client.DoAction(&req, &resp)
	return
}

type SubscribeImageResponse struct {
	RequestId string
	Success   bool
}
