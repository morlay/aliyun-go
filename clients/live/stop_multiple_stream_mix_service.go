package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type StopMultipleStreamMixServiceRequest struct {
	requests.RpcRequest
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
}

func (req *StopMultipleStreamMixServiceRequest) Invoke(client *sdk.Client) (resp *StopMultipleStreamMixServiceResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "StopMultipleStreamMixService", "", "")
	resp = &StopMultipleStreamMixServiceResponse{}
	err = client.DoAction(req, resp)
	return
}

type StopMultipleStreamMixServiceResponse struct {
	responses.BaseResponse
	RequestId string
}
