package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeLiveDetectNotifyConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeLiveDetectNotifyConfigRequest) Invoke(client *sdk.Client) (response *DescribeLiveDetectNotifyConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeLiveDetectNotifyConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "DescribeLiveDetectNotifyConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeLiveDetectNotifyConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeLiveDetectNotifyConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeLiveDetectNotifyConfigResponse struct {
	RequestId              string
	LiveDetectNotifyConfig DescribeLiveDetectNotifyConfigLiveDetectNotifyConfig
}

type DescribeLiveDetectNotifyConfigLiveDetectNotifyConfig struct {
	DomainName string
	NotifyUrl  string
}
