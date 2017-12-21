package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteLiveStreamsNotifyUrlConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DeleteLiveStreamsNotifyUrlConfigRequest) Invoke(client *sdk.Client) (response *DeleteLiveStreamsNotifyUrlConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteLiveStreamsNotifyUrlConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "DeleteLiveStreamsNotifyUrlConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteLiveStreamsNotifyUrlConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteLiveStreamsNotifyUrlConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteLiveStreamsNotifyUrlConfigResponse struct {
	RequestId string
}
