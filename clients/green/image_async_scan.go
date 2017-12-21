package green

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ImageAsyncScanRequest struct {
	requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

func (req *ImageAsyncScanRequest) Invoke(client *sdk.Client) (resp *ImageAsyncScanResponse, err error) {
	req.InitWithApiInfo("Green", "2017-01-12", "ImageAsyncScan", "/green/image/asyncscan", "green", "")
	req.Method = "POST"

	resp = &ImageAsyncScanResponse{}
	err = client.DoAction(req, resp)
	return
}

type ImageAsyncScanResponse struct {
	responses.BaseResponse
}
