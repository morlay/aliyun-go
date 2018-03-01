package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteLiveStreamsNotifyUrlConfigRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteLiveStreamsNotifyUrlConfigRequest) Invoke(client *sdk.Client) (resp *DeleteLiveStreamsNotifyUrlConfigResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DeleteLiveStreamsNotifyUrlConfig", "live", "")
	resp = &DeleteLiveStreamsNotifyUrlConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteLiveStreamsNotifyUrlConfigResponse struct {
	responses.BaseResponse
	RequestId string
}
