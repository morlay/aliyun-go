package green

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ImageAsyncScanResultsRequest struct {
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

func (r ImageAsyncScanResultsRequest) Invoke(client *sdk.Client) (response *ImageAsyncScanResultsResponse, err error) {
	req := struct {
		*requests.RoaRequest
		ImageAsyncScanResultsRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("Green", "2017-01-12", "ImageAsyncScanResults", "/green/image/results", "green", "")
	req.Method = "POST"

	resp := struct {
		*responses.BaseResponse
		ImageAsyncScanResultsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ImageAsyncScanResultsResponse

	err = client.DoAction(&req, &resp)
	return
}

type ImageAsyncScanResultsResponse struct {
}
