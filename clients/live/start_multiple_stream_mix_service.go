package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type StartMultipleStreamMixServiceRequest struct {
	requests.RpcRequest
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	MixTemplate   string `position:"Query" name:"MixTemplate"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
}

func (req *StartMultipleStreamMixServiceRequest) Invoke(client *sdk.Client) (resp *StartMultipleStreamMixServiceResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "StartMultipleStreamMixService", "", "")
	resp = &StartMultipleStreamMixServiceResponse{}
	err = client.DoAction(req, resp)
	return
}

type StartMultipleStreamMixServiceResponse struct {
	responses.BaseResponse
	RequestId string
}
