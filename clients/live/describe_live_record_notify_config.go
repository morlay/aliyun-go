package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveRecordNotifyConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeLiveRecordNotifyConfigRequest) Invoke(client *sdk.Client) (response *DescribeLiveRecordNotifyConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeLiveRecordNotifyConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "DescribeLiveRecordNotifyConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeLiveRecordNotifyConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeLiveRecordNotifyConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeLiveRecordNotifyConfigResponse struct {
	RequestId              string
	LiveRecordNotifyConfig DescribeLiveRecordNotifyConfigLiveRecordNotifyConfig
}

type DescribeLiveRecordNotifyConfigLiveRecordNotifyConfig struct {
	DomainName       string
	NotifyUrl        string
	NeedStatusNotify bool
}
