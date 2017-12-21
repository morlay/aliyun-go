package qualitycheck

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UploadAudioDataWithRules4PreRequest struct {
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (r UploadAudioDataWithRules4PreRequest) Invoke(client *sdk.Client) (response *UploadAudioDataWithRules4PreResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UploadAudioDataWithRules4PreRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "UploadAudioDataWithRules4Pre", "", "")

	resp := struct {
		*responses.BaseResponse
		UploadAudioDataWithRules4PreResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UploadAudioDataWithRules4PreResponse

	err = client.DoAction(&req, &resp)
	return
}

type UploadAudioDataWithRules4PreResponse struct {
	RequestId string
	Success   bool
	Code      string
	Message   string
	Data      string
}
