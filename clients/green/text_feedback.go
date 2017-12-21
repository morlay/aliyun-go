package green

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type TextFeedbackRequest struct {
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

func (r TextFeedbackRequest) Invoke(client *sdk.Client) (response *TextFeedbackResponse, err error) {
	req := struct {
		*requests.RoaRequest
		TextFeedbackRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("Green", "2017-01-12", "TextFeedback", "/green/text/feedback", "green", "")
	req.Method = "POST"

	resp := struct {
		*responses.BaseResponse
		TextFeedbackResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.TextFeedbackResponse

	err = client.DoAction(&req, &resp)
	return
}

type TextFeedbackResponse struct {
}
