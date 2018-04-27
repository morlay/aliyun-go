package green

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ImageSyncScanRequest struct {
	requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

func (req *ImageSyncScanRequest) Invoke(client *sdk.Client) (resp *ImageSyncScanResponse, err error) {
	req.InitWithApiInfo("Green", "2017-08-25", "ImageSyncScan", "/green/image/scan", "green", "")
	req.Method = "POST"

	resp = &ImageSyncScanResponse{}
	err = client.DoAction(req, resp)
	return
}

type ImageSyncScanResponse struct {
	responses.BaseResponse
}
