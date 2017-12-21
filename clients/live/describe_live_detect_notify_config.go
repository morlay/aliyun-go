package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveDetectNotifyConfigRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeLiveDetectNotifyConfigRequest) Invoke(client *sdk.Client) (resp *DescribeLiveDetectNotifyConfigResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DescribeLiveDetectNotifyConfig", "", "")
	resp = &DescribeLiveDetectNotifyConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveDetectNotifyConfigResponse struct {
	responses.BaseResponse
	RequestId              string
	LiveDetectNotifyConfig DescribeLiveDetectNotifyConfigLiveDetectNotifyConfig
}

type DescribeLiveDetectNotifyConfigLiveDetectNotifyConfig struct {
	DomainName string
	NotifyUrl  string
}
