package qualitycheck

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DelThesaurusForApiRequest struct {
	requests.RpcRequest
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (req *DelThesaurusForApiRequest) Invoke(client *sdk.Client) (resp *DelThesaurusForApiResponse, err error) {
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "DelThesaurusForApi", "", "")
	resp = &DelThesaurusForApiResponse{}
	err = client.DoAction(req, resp)
	return
}

type DelThesaurusForApiResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Code      common.String
	Message   common.String
}
