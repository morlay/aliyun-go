package qualitycheck

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UploadAudioData4PreRequest struct {
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (r UploadAudioData4PreRequest) Invoke(client *sdk.Client) (response *UploadAudioData4PreResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UploadAudioData4PreRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "UploadAudioData4Pre", "", "")

	resp := struct {
		*responses.BaseResponse
		UploadAudioData4PreResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UploadAudioData4PreResponse

	err = client.DoAction(&req, &resp)
	return
}

type UploadAudioData4PreResponse struct {
	RequestId string
	Success   bool
	Code      string
	Message   string
	Data      string
}
