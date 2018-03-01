package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveStreamsNotifyUrlConfigRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeLiveStreamsNotifyUrlConfigRequest) Invoke(client *sdk.Client) (resp *DescribeLiveStreamsNotifyUrlConfigResponse, err error) {
	req.InitWithApiInfo("live", "2016-11-01", "DescribeLiveStreamsNotifyUrlConfig", "live", "")
	resp = &DescribeLiveStreamsNotifyUrlConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeLiveStreamsNotifyUrlConfigResponse struct {
	responses.BaseResponse
	RequestId               string
	LiveStreamsNotifyConfig DescribeLiveStreamsNotifyUrlConfigLiveStreamsNotifyConfig
}

type DescribeLiveStreamsNotifyUrlConfigLiveStreamsNotifyConfig struct {
	DomainName string
	NotifyUrl  string
}
