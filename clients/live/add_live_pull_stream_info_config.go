package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddLivePullStreamInfoConfigRequest struct {
	SourceUrl     string `position:"Query" name:"SourceUrl"`
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
}

func (r AddLivePullStreamInfoConfigRequest) Invoke(client *sdk.Client) (response *AddLivePullStreamInfoConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddLivePullStreamInfoConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "AddLivePullStreamInfoConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		AddLivePullStreamInfoConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddLivePullStreamInfoConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddLivePullStreamInfoConfigResponse struct {
	RequestId string
}
