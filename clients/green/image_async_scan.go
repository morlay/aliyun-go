package green

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ImageAsyncScanRequest struct {
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

func (r ImageAsyncScanRequest) Invoke(client *sdk.Client) (response *ImageAsyncScanResponse, err error) {
	req := struct {
		*requests.RoaRequest
		ImageAsyncScanRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("Green", "2017-01-12", "ImageAsyncScan", "/green/image/asyncscan", "green", "")
	req.Method = "POST"

	resp := struct {
		*responses.BaseResponse
		ImageAsyncScanResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ImageAsyncScanResponse

	err = client.DoAction(&req, &resp)
	return
}

type ImageAsyncScanResponse struct {
}
