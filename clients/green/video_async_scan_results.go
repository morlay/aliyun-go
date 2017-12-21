package green

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type VideoAsyncScanResultsRequest struct {
	requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

func (req *VideoAsyncScanResultsRequest) Invoke(client *sdk.Client) (resp *VideoAsyncScanResultsResponse, err error) {
	req.InitWithApiInfo("Green", "2017-01-12", "VideoAsyncScanResults", "/green/video/results", "green", "")
	req.Method = "POST"

	resp = &VideoAsyncScanResultsResponse{}
	err = client.DoAction(req, resp)
	return
}

type VideoAsyncScanResultsResponse struct {
	responses.BaseResponse
}
