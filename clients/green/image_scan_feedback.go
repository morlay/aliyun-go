package green

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ImageScanFeedbackRequest struct {
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

func (r ImageScanFeedbackRequest) Invoke(client *sdk.Client) (response *ImageScanFeedbackResponse, err error) {
	req := struct {
		*requests.RoaRequest
		ImageScanFeedbackRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("Green", "2017-01-12", "ImageScanFeedback", "/green/image/feedback", "green", "")
	req.Method = "POST"

	resp := struct {
		*responses.BaseResponse
		ImageScanFeedbackResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ImageScanFeedbackResponse

	err = client.DoAction(&req, &resp)
	return
}

type ImageScanFeedbackResponse struct {
}
