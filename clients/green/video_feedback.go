package green

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type VideoFeedbackRequest struct {
	requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

func (req *VideoFeedbackRequest) Invoke(client *sdk.Client) (resp *VideoFeedbackResponse, err error) {
	req.InitWithApiInfo("Green", "2017-08-25", "VideoFeedback", "/green/video/feedback", "green", "")
	req.Method = "POST"

	resp = &VideoFeedbackResponse{}
	err = client.DoAction(req, resp)
	return
}

type VideoFeedbackResponse struct {
	responses.BaseResponse
}
