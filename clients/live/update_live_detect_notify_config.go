package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateLiveDetectNotifyConfigRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	NotifyUrl     string `position:"Query" name:"NotifyUrl"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *UpdateLiveDetectNotifyConfigRequest) Invoke(client *sdk.Client) (resp *UpdateLiveDetectNotifyConfigResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "UpdateLiveDetectNotifyConfig", "live", "")
	resp = &UpdateLiveDetectNotifyConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateLiveDetectNotifyConfigResponse struct {
	responses.BaseResponse
	RequestId string
}
