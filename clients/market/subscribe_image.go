package market

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SubscribeImageRequest struct {
	requests.RpcRequest
	ProductCode string `position:"Query" name:"ProductCode"`
}

func (req *SubscribeImageRequest) Invoke(client *sdk.Client) (resp *SubscribeImageResponse, err error) {
	req.InitWithApiInfo("Market", "2015-11-01", "SubscribeImage", "yunmarket", "")
	resp = &SubscribeImageResponse{}
	err = client.DoAction(req, resp)
	return
}

type SubscribeImageResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
}
