package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AddCustomLiveStreamTranscodeRequest struct {
	requests.RpcRequest
	App           string `position:"Query" name:"App"`
	Template      string `position:"Query" name:"Template"`
	FPS           int    `position:"Query" name:"FPS"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	TemplateType  string `position:"Query" name:"TemplateType"`
	Domain        string `position:"Query" name:"Domain"`
	Width         int    `position:"Query" name:"Width"`
	VideoBitrate  int    `position:"Query" name:"VideoBitrate"`
	Height        int    `position:"Query" name:"Height"`
}

func (req *AddCustomLiveStreamTranscodeRequest) Invoke(client *sdk.Client) (resp *AddCustomLiveStreamTranscodeResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "AddCustomLiveStreamTranscode", "live", "")
	resp = &AddCustomLiveStreamTranscodeResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddCustomLiveStreamTranscodeResponse struct {
	responses.BaseResponse
	RequestId common.String
}
