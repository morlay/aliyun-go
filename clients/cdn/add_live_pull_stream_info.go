package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddLivePullStreamInfoRequest struct {
	requests.RpcRequest
	SourceUrl     string `position:"Query" name:"SourceUrl"`
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
}

func (req *AddLivePullStreamInfoRequest) Invoke(client *sdk.Client) (resp *AddLivePullStreamInfoResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "AddLivePullStreamInfo", "", "")
	resp = &AddLivePullStreamInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddLivePullStreamInfoResponse struct {
	responses.BaseResponse
	RequestId string
}
