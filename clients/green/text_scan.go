package green

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type TextScanRequest struct {
	requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

func (req *TextScanRequest) Invoke(client *sdk.Client) (resp *TextScanResponse, err error) {
	req.InitWithApiInfo("Green", "2017-08-25", "TextScan", "/green/text/scan", "green", "")
	req.Method = "POST"

	resp = &TextScanResponse{}
	err = client.DoAction(req, resp)
	return
}

type TextScanResponse struct {
	responses.BaseResponse
}
