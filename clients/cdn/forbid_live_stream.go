package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ForbidLiveStreamRequest struct {
	requests.RpcRequest
	ResumeTime     string `position:"Query" name:"ResumeTime"`
	AppName        string `position:"Query" name:"AppName"`
	SecurityToken  string `position:"Query" name:"SecurityToken"`
	LiveStreamType string `position:"Query" name:"LiveStreamType"`
	DomainName     string `position:"Query" name:"DomainName"`
	OwnerId        int64  `position:"Query" name:"OwnerId"`
	StreamName     string `position:"Query" name:"StreamName"`
}

func (req *ForbidLiveStreamRequest) Invoke(client *sdk.Client) (resp *ForbidLiveStreamResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "ForbidLiveStream", "", "")
	resp = &ForbidLiveStreamResponse{}
	err = client.DoAction(req, resp)
	return
}

type ForbidLiveStreamResponse struct {
	responses.BaseResponse
	RequestId string
}
