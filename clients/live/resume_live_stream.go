package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ResumeLiveStreamRequest struct {
	requests.RpcRequest
	AppName        string `position:"Query" name:"AppName"`
	SecurityToken  string `position:"Query" name:"SecurityToken"`
	LiveStreamType string `position:"Query" name:"LiveStreamType"`
	DomainName     string `position:"Query" name:"DomainName"`
	OwnerId        int64  `position:"Query" name:"OwnerId"`
	StreamName     string `position:"Query" name:"StreamName"`
}

func (req *ResumeLiveStreamRequest) Invoke(client *sdk.Client) (resp *ResumeLiveStreamResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "ResumeLiveStream", "", "")
	resp = &ResumeLiveStreamResponse{}
	err = client.DoAction(req, resp)
	return
}

type ResumeLiveStreamResponse struct {
	responses.BaseResponse
	RequestId string
}
