package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RemoveMultipleStreamMixServiceRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	MixStreamName string `position:"Query" name:"MixStreamName"`
	MixDomainName string `position:"Query" name:"MixDomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	MixAppName    string `position:"Query" name:"MixAppName"`
	StreamName    string `position:"Query" name:"StreamName"`
}

func (r RemoveMultipleStreamMixServiceRequest) Invoke(client *sdk.Client) (response *RemoveMultipleStreamMixServiceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		RemoveMultipleStreamMixServiceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "RemoveMultipleStreamMixService", "", "")

	resp := struct {
		*responses.BaseResponse
		RemoveMultipleStreamMixServiceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.RemoveMultipleStreamMixServiceResponse

	err = client.DoAction(&req, &resp)
	return
}

type RemoveMultipleStreamMixServiceResponse struct {
	RequestId string
}
