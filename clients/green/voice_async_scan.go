package green

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type VoiceAsyncScanRequest struct {
	requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

func (req *VoiceAsyncScanRequest) Invoke(client *sdk.Client) (resp *VoiceAsyncScanResponse, err error) {
	req.InitWithApiInfo("Green", "2017-08-25", "VoiceAsyncScan", "/green/voice/asyncscan", "green", "")
	req.Method = "POST"

	resp = &VoiceAsyncScanResponse{}
	err = client.DoAction(req, resp)
	return
}

type VoiceAsyncScanResponse struct {
	responses.BaseResponse
}
