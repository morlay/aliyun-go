package qualitycheck

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UploadAudioData4PreRequest struct {
	requests.RpcRequest
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (req *UploadAudioData4PreRequest) Invoke(client *sdk.Client) (resp *UploadAudioData4PreResponse, err error) {
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "UploadAudioData4Pre", "", "")
	resp = &UploadAudioData4PreResponse{}
	err = client.DoAction(req, resp)
	return
}

type UploadAudioData4PreResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Code      string
	Message   string
	Data      string
}
