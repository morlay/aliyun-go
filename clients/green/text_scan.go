package green

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type TextScanRequest struct {
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

func (r TextScanRequest) Invoke(client *sdk.Client) (response *TextScanResponse, err error) {
	req := struct {
		*requests.RoaRequest
		TextScanRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("Green", "2017-01-12", "TextScan", "/green/text/scan", "green", "")
	req.Method = "POST"

	resp := struct {
		*responses.BaseResponse
		TextScanResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.TextScanResponse

	err = client.DoAction(&req, &resp)
	return
}

type TextScanResponse struct {
}
