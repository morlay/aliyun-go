package qualitycheck

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type UpdateScoreForApiRequest struct {
	requests.RpcRequest
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (req *UpdateScoreForApiRequest) Invoke(client *sdk.Client) (resp *UpdateScoreForApiResponse, err error) {
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "UpdateScoreForApi", "", "")
	resp = &UpdateScoreForApiResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateScoreForApiResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Code      common.String
	Message   common.String
}
