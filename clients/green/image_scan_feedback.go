package green

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ImageScanFeedbackRequest struct {
	requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

func (req *ImageScanFeedbackRequest) Invoke(client *sdk.Client) (resp *ImageScanFeedbackResponse, err error) {
	req.InitWithApiInfo("Green", "2017-08-25", "ImageScanFeedback", "/green/image/feedback", "green", "")
	req.Method = "POST"

	resp = &ImageScanFeedbackResponse{}
	err = client.DoAction(req, resp)
	return
}

type ImageScanFeedbackResponse struct {
	responses.BaseResponse
}
