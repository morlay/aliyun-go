package green

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type VoiceAsyncScanResultsRequest struct {
	requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

func (req *VoiceAsyncScanResultsRequest) Invoke(client *sdk.Client) (resp *VoiceAsyncScanResultsResponse, err error) {
	req.InitWithApiInfo("Green", "2017-08-25", "VoiceAsyncScanResults", "/green/voice/results", "green", "")
	req.Method = "POST"

	resp = &VoiceAsyncScanResultsResponse{}
	err = client.DoAction(req, resp)
	return
}

type VoiceAsyncScanResultsResponse struct {
	responses.BaseResponse
}
