package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteLiveStreamTranscodeRequest struct {
	App           string `position:"Query" name:"App"`
	Template      string `position:"Query" name:"Template"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	Domain        string `position:"Query" name:"Domain"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DeleteLiveStreamTranscodeRequest) Invoke(client *sdk.Client) (response *DeleteLiveStreamTranscodeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteLiveStreamTranscodeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "DeleteLiveStreamTranscode", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteLiveStreamTranscodeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteLiveStreamTranscodeResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteLiveStreamTranscodeResponse struct {
	RequestId string
}
