package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type RemoveMultipleStreamMixServiceRequest struct {
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

func (req *RemoveMultipleStreamMixServiceRequest) Invoke(client *sdk.Client) (resp *RemoveMultipleStreamMixServiceResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "RemoveMultipleStreamMixService", "live", "")
	resp = &RemoveMultipleStreamMixServiceResponse{}
	err = client.DoAction(req, resp)
	return
}

type RemoveMultipleStreamMixServiceResponse struct {
	responses.BaseResponse
	RequestId common.String
}
