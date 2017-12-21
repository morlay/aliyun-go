package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteLiveStreamTranscodeRequest struct {
	requests.RpcRequest
	Template      string `position:"Query" name:"Template"`
	App           string `position:"Query" name:"App"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerAccount  string `position:"Query" name:"OwnerAccount"`
	Domain        string `position:"Query" name:"Domain"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteLiveStreamTranscodeRequest) Invoke(client *sdk.Client) (resp *DeleteLiveStreamTranscodeResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DeleteLiveStreamTranscode", "", "")
	resp = &DeleteLiveStreamTranscodeResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteLiveStreamTranscodeResponse struct {
	responses.BaseResponse
	RequestId string
}
