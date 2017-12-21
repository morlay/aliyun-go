package green

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type VideoAsyncScanRequest struct {
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

func (r VideoAsyncScanRequest) Invoke(client *sdk.Client) (response *VideoAsyncScanResponse, err error) {
	req := struct {
		*requests.RoaRequest
		VideoAsyncScanRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("Green", "2017-01-12", "VideoAsyncScan", "/green/video/asyncscan", "green", "")
	req.Method = "POST"

	resp := struct {
		*responses.BaseResponse
		VideoAsyncScanResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.VideoAsyncScanResponse

	err = client.DoAction(&req, &resp)
	return
}

type VideoAsyncScanResponse struct {
}
