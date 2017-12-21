package qualitycheck

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveReviewResultRequest struct {
	requests.RpcRequest
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (req *SaveReviewResultRequest) Invoke(client *sdk.Client) (resp *SaveReviewResultResponse, err error) {
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "SaveReviewResult", "", "")
	resp = &SaveReviewResultResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveReviewResultResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Code      string
	Message   string
	Data      string
}
