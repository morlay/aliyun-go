package qualitycheck

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId common.String
	Success   bool
	Code      common.String
	Message   common.String
	Data      common.String
}
