package qualitycheck

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteSubScoreForApiRequest struct {
	requests.RpcRequest
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (req *DeleteSubScoreForApiRequest) Invoke(client *sdk.Client) (resp *DeleteSubScoreForApiResponse, err error) {
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "DeleteSubScoreForApi", "", "")
	resp = &DeleteSubScoreForApiResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteSubScoreForApiResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Code      common.String
	Message   common.String
}
