package green

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type TextFeedbackRequest struct {
	requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

func (req *TextFeedbackRequest) Invoke(client *sdk.Client) (resp *TextFeedbackResponse, err error) {
	req.InitWithApiInfo("Green", "2017-08-25", "TextFeedback", "/green/text/feedback", "green", "")
	req.Method = "POST"

	resp = &TextFeedbackResponse{}
	err = client.DoAction(req, resp)
	return
}

type TextFeedbackResponse struct {
	responses.BaseResponse
}
