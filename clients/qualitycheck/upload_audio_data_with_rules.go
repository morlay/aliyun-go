package qualitycheck

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UploadAudioDataWithRulesRequest struct {
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (r UploadAudioDataWithRulesRequest) Invoke(client *sdk.Client) (response *UploadAudioDataWithRulesResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UploadAudioDataWithRulesRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "UploadAudioDataWithRules", "", "")

	resp := struct {
		*responses.BaseResponse
		UploadAudioDataWithRulesResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UploadAudioDataWithRulesResponse

	err = client.DoAction(&req, &resp)
	return
}

type UploadAudioDataWithRulesResponse struct {
	RequestId string
	Success   bool
	Code      string
	Message   string
	Data      string
}
