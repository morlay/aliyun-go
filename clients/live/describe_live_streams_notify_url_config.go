package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveStreamsNotifyUrlConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeLiveStreamsNotifyUrlConfigRequest) Invoke(client *sdk.Client) (response *DescribeLiveStreamsNotifyUrlConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeLiveStreamsNotifyUrlConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "DescribeLiveStreamsNotifyUrlConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeLiveStreamsNotifyUrlConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeLiveStreamsNotifyUrlConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeLiveStreamsNotifyUrlConfigResponse struct {
	RequestId               string
	LiveStreamsNotifyConfig DescribeLiveStreamsNotifyUrlConfigLiveStreamsNotifyConfig
}

type DescribeLiveStreamsNotifyUrlConfigLiveStreamsNotifyConfig struct {
	DomainName string
	NotifyUrl  string
}
