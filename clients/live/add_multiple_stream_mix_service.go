package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddMultipleStreamMixServiceRequest struct {
	requests.RpcRequest
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	MixStreamName string `position:"Query" name:"MixStreamName"`
	MixDomainName string `position:"Query" name:"MixDomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	MixAppName    string `position:"Query" name:"MixAppName"`
	StreamName    string `position:"Query" name:"StreamName"`
}

func (req *AddMultipleStreamMixServiceRequest) Invoke(client *sdk.Client) (resp *AddMultipleStreamMixServiceResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "AddMultipleStreamMixService", "", "")
	resp = &AddMultipleStreamMixServiceResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddMultipleStreamMixServiceResponse struct {
	responses.BaseResponse
	RequestId string
}
