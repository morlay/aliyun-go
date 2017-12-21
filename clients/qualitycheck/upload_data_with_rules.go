package qualitycheck

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UploadDataWithRulesRequest struct {
	requests.RpcRequest
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (req *UploadDataWithRulesRequest) Invoke(client *sdk.Client) (resp *UploadDataWithRulesResponse, err error) {
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "UploadDataWithRules", "", "")
	resp = &UploadDataWithRulesResponse{}
	err = client.DoAction(req, resp)
	return
}

type UploadDataWithRulesResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Code      string
	Message   string
	Data      string
}
