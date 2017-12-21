package green

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type VideoAsyncScanResultsRequest struct {
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

func (r VideoAsyncScanResultsRequest) Invoke(client *sdk.Client) (response *VideoAsyncScanResultsResponse, err error) {
	req := struct {
		*requests.RoaRequest
		VideoAsyncScanResultsRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("Green", "2017-01-12", "VideoAsyncScanResults", "/green/video/results", "green", "")
	req.Method = "POST"

	resp := struct {
		*responses.BaseResponse
		VideoAsyncScanResultsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.VideoAsyncScanResultsResponse

	err = client.DoAction(&req, &resp)
	return
}

type VideoAsyncScanResultsResponse struct {
}
