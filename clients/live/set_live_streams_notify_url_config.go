package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetLiveStreamsNotifyUrlConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	NotifyUrl     string `position:"Query" name:"NotifyUrl"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r SetLiveStreamsNotifyUrlConfigRequest) Invoke(client *sdk.Client) (response *SetLiveStreamsNotifyUrlConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetLiveStreamsNotifyUrlConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "SetLiveStreamsNotifyUrlConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		SetLiveStreamsNotifyUrlConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetLiveStreamsNotifyUrlConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetLiveStreamsNotifyUrlConfigResponse struct {
	RequestId string
}
