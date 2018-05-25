package qualitycheck

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AddThesaurusForApiRequest struct {
	requests.RpcRequest
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (req *AddThesaurusForApiRequest) Invoke(client *sdk.Client) (resp *AddThesaurusForApiResponse, err error) {
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "AddThesaurusForApi", "", "")
	resp = &AddThesaurusForApiResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddThesaurusForApiResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Code      common.String
	Message   common.String
	Data      common.Long
}
