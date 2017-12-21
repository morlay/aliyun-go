package qualitycheck

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UploadAudioDataWithRulesRequest struct {
	requests.RpcRequest
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (req *UploadAudioDataWithRulesRequest) Invoke(client *sdk.Client) (resp *UploadAudioDataWithRulesResponse, err error) {
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "UploadAudioDataWithRules", "", "")
	resp = &UploadAudioDataWithRulesResponse{}
	err = client.DoAction(req, resp)
	return
}

type UploadAudioDataWithRulesResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Code      string
	Message   string
	Data      string
}
