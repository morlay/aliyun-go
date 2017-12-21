package green

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ImageSyncScanRequest struct {
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

func (r ImageSyncScanRequest) Invoke(client *sdk.Client) (response *ImageSyncScanResponse, err error) {
	req := struct {
		*requests.RoaRequest
		ImageSyncScanRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("Green", "2017-01-12", "ImageSyncScan", "/green/image/scan", "green", "")
	req.Method = "POST"

	resp := struct {
		*responses.BaseResponse
		ImageSyncScanResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ImageSyncScanResponse

	err = client.DoAction(&req, &resp)
	return
}

type ImageSyncScanResponse struct {
}
