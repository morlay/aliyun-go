package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddLiveDetectNotifyConfigRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	NotifyUrl     string `position:"Query" name:"NotifyUrl"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *AddLiveDetectNotifyConfigRequest) Invoke(client *sdk.Client) (resp *AddLiveDetectNotifyConfigResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "AddLiveDetectNotifyConfig", "", "")
	resp = &AddLiveDetectNotifyConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddLiveDetectNotifyConfigResponse struct {
	responses.BaseResponse
	RequestId string
}
