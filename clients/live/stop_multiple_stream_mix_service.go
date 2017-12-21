package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type StopMultipleStreamMixServiceRequest struct {
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
}

func (r StopMultipleStreamMixServiceRequest) Invoke(client *sdk.Client) (response *StopMultipleStreamMixServiceResponse, err error) {
	req := struct {
		*requests.RpcRequest
		StopMultipleStreamMixServiceRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "StopMultipleStreamMixService", "", "")

	resp := struct {
		*responses.BaseResponse
		StopMultipleStreamMixServiceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.StopMultipleStreamMixServiceResponse

	err = client.DoAction(&req, &resp)
	return
}

type StopMultipleStreamMixServiceResponse struct {
	RequestId string
}
