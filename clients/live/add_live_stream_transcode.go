package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AddLiveStreamTranscodeRequest struct {
	requests.RpcRequest
	App           string `position:"Query" name:"App"`
	Template      string `position:"Query" name:"Template"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	Domain        string `position:"Query" name:"Domain"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *AddLiveStreamTranscodeRequest) Invoke(client *sdk.Client) (resp *AddLiveStreamTranscodeResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "AddLiveStreamTranscode", "live", "")
	resp = &AddLiveStreamTranscodeResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddLiveStreamTranscodeResponse struct {
	responses.BaseResponse
	RequestId common.String
}
