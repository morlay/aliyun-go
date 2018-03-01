package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteLiveStreamTranscodeRequest struct {
	requests.RpcRequest
	App           string `position:"Query" name:"App"`
	Template      string `position:"Query" name:"Template"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	Domain        string `position:"Query" name:"Domain"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteLiveStreamTranscodeRequest) Invoke(client *sdk.Client) (resp *DeleteLiveStreamTranscodeResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DeleteLiveStreamTranscode", "live", "")
	resp = &DeleteLiveStreamTranscodeResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteLiveStreamTranscodeResponse struct {
	responses.BaseResponse
	RequestId string
}
