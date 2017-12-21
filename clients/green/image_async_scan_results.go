package green

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ImageAsyncScanResultsRequest struct {
	requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

func (req *ImageAsyncScanResultsRequest) Invoke(client *sdk.Client) (resp *ImageAsyncScanResultsResponse, err error) {
	req.InitWithApiInfo("Green", "2017-01-12", "ImageAsyncScanResults", "/green/image/results", "green", "")
	req.Method = "POST"

	resp = &ImageAsyncScanResultsResponse{}
	err = client.DoAction(req, resp)
	return
}

type ImageAsyncScanResultsResponse struct {
	responses.BaseResponse
}
