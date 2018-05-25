package qualitycheck

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteScoreForApiRequest struct {
	requests.RpcRequest
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (req *DeleteScoreForApiRequest) Invoke(client *sdk.Client) (resp *DeleteScoreForApiResponse, err error) {
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "DeleteScoreForApi", "", "")
	resp = &DeleteScoreForApiResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteScoreForApiResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Code      common.String
	Message   common.String
}
