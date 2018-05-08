package qualitycheck

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId common.String
	Success   bool
	Code      common.String
	Message   common.String
	Data      common.String
}
