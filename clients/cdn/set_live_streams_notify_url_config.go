package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SetLiveStreamsNotifyUrlConfigRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	NotifyUrl     string `position:"Query" name:"NotifyUrl"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *SetLiveStreamsNotifyUrlConfigRequest) Invoke(client *sdk.Client) (resp *SetLiveStreamsNotifyUrlConfigResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetLiveStreamsNotifyUrlConfig", "", "")
	resp = &SetLiveStreamsNotifyUrlConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetLiveStreamsNotifyUrlConfigResponse struct {
	responses.BaseResponse
	RequestId common.String
}
