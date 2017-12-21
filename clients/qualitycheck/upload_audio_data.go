package qualitycheck

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UploadAudioDataRequest struct {
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (r UploadAudioDataRequest) Invoke(client *sdk.Client) (response *UploadAudioDataResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UploadAudioDataRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "UploadAudioData", "", "")

	resp := struct {
		*responses.BaseResponse
		UploadAudioDataResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UploadAudioDataResponse

	err = client.DoAction(&req, &resp)
	return
}

type UploadAudioDataResponse struct {
	RequestId string
	Success   bool
	Code      string
	Message   string
	Data      string
}
