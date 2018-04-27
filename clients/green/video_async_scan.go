package green

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type VideoAsyncScanRequest struct {
	requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

func (req *VideoAsyncScanRequest) Invoke(client *sdk.Client) (resp *VideoAsyncScanResponse, err error) {
	req.InitWithApiInfo("Green", "2017-08-25", "VideoAsyncScan", "/green/video/asyncscan", "green", "")
	req.Method = "POST"

	resp = &VideoAsyncScanResponse{}
	err = client.DoAction(req, resp)
	return
}

type VideoAsyncScanResponse struct {
	responses.BaseResponse
}
