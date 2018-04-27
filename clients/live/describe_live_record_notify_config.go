package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveRecordNotifyConfigRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeLiveRecordNotifyConfigRequest) Invoke(client *sdk.Client) (resp *DescribeLiveRecordNotifyConfigResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DescribeLiveRecordNotifyConfig", "live", "")
	resp = &DescribeLiveRecordNotifyConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveRecordNotifyConfigResponse struct {
	responses.BaseResponse
	RequestId              string
	LiveRecordNotifyConfig DescribeLiveRecordNotifyConfigLiveRecordNotifyConfig
}

type DescribeLiveRecordNotifyConfigLiveRecordNotifyConfig struct {
	DomainName       string
	NotifyUrl        string
	OnDemandUrl      string
	NeedStatusNotify bool
}
