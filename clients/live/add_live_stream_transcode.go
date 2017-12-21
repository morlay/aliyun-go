package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddLiveStreamTranscodeRequest struct {
	App           string `position:"Query" name:"App"`
	Template      string `position:"Query" name:"Template"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	Domain        string `position:"Query" name:"Domain"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r AddLiveStreamTranscodeRequest) Invoke(client *sdk.Client) (response *AddLiveStreamTranscodeResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddLiveStreamTranscodeRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "AddLiveStreamTranscode", "", "")

	resp := struct {
		*responses.BaseResponse
		AddLiveStreamTranscodeResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddLiveStreamTranscodeResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddLiveStreamTranscodeResponse struct {
	RequestId string
}
