package qualitycheck

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveReviewResultRequest struct {
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (r SaveReviewResultRequest) Invoke(client *sdk.Client) (response *SaveReviewResultResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SaveReviewResultRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "SaveReviewResult", "", "")

	resp := struct {
		*responses.BaseResponse
		SaveReviewResultResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SaveReviewResultResponse

	err = client.DoAction(&req, &resp)
	return
}

type SaveReviewResultResponse struct {
	RequestId string
	Success   bool
	Code      string
	Message   string
	Data      string
}
