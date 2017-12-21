package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteLiveStreamTranscodeRequest struct {
	Template      string `position:"Query" name:"Template"`
	App           string `position:"Query" name:"App"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerAccount  string `position:"Query" name:"OwnerAccount"`
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
	req.InitWithApiInfo("Cdn", "2014-11-11", "DeleteLiveStreamTranscode", "", "")

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
