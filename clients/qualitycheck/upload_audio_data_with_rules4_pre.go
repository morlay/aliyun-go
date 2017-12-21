package qualitycheck

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UploadAudioDataWithRules4PreRequest struct {
	requests.RpcRequest
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (req *UploadAudioDataWithRules4PreRequest) Invoke(client *sdk.Client) (resp *UploadAudioDataWithRules4PreResponse, err error) {
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "UploadAudioDataWithRules4Pre", "", "")
	resp = &UploadAudioDataWithRules4PreResponse{}
	err = client.DoAction(req, resp)
	return
}

type UploadAudioDataWithRules4PreResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Code      string
	Message   string
	Data      string
}
