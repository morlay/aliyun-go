package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateLiveRecordNotifyConfigRequest struct {
	requests.RpcRequest
	SecurityToken    string `position:"Query" name:"SecurityToken"`
	DomainName       string `position:"Query" name:"DomainName"`
	NotifyUrl        string `position:"Query" name:"NotifyUrl"`
	OwnerId          int64  `position:"Query" name:"OwnerId"`
	NeedStatusNotify string `position:"Query" name:"NeedStatusNotify"`
}

func (req *UpdateLiveRecordNotifyConfigRequest) Invoke(client *sdk.Client) (resp *UpdateLiveRecordNotifyConfigResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "UpdateLiveRecordNotifyConfig", "live", "")
	resp = &UpdateLiveRecordNotifyConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateLiveRecordNotifyConfigResponse struct {
	responses.BaseResponse
	RequestId string
}
