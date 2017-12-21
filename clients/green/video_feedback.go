package green

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type VideoFeedbackRequest struct {
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

func (r VideoFeedbackRequest) Invoke(client *sdk.Client) (response *VideoFeedbackResponse, err error) {
	req := struct {
		*requests.RoaRequest
		VideoFeedbackRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("Green", "2017-01-12", "VideoFeedback", "/green/video/feedback", "green", "")
	req.Method = "POST"

	resp := struct {
		*responses.BaseResponse
		VideoFeedbackResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.VideoFeedbackResponse

	err = client.DoAction(&req, &resp)
	return
}

type VideoFeedbackResponse struct {
}
