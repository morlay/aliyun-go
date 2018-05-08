package qualitycheck

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type UploadAudioDataRequest struct {
	requests.RpcRequest
	JsonStr string `position:"Query" name:"JsonStr"`
}

func (req *UploadAudioDataRequest) Invoke(client *sdk.Client) (resp *UploadAudioDataResponse, err error) {
	req.InitWithApiInfo("Qualitycheck", "2016-08-01", "UploadAudioData", "", "")
	resp = &UploadAudioDataResponse{}
	err = client.DoAction(req, resp)
	return
}

type UploadAudioDataResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Code      common.String
	Message   common.String
	Data      common.String
}
