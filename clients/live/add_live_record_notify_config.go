package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddLiveRecordNotifyConfigRequest struct {
	requests.RpcRequest
	SecurityToken    string `position:"Query" name:"SecurityToken"`
	DomainName       string `position:"Query" name:"DomainName"`
	NotifyUrl        string `position:"Query" name:"NotifyUrl"`
	OwnerId          int64  `position:"Query" name:"OwnerId"`
	NeedStatusNotify string `position:"Query" name:"NeedStatusNotify"`
}

func (req *AddLiveRecordNotifyConfigRequest) Invoke(client *sdk.Client) (resp *AddLiveRecordNotifyConfigResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "AddLiveRecordNotifyConfig", "live", "")
	resp = &AddLiveRecordNotifyConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddLiveRecordNotifyConfigResponse struct {
	responses.BaseResponse
	RequestId string
}
