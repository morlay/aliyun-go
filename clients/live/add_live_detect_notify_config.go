package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddLiveDetectNotifyConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	NotifyUrl     string `position:"Query" name:"NotifyUrl"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r AddLiveDetectNotifyConfigRequest) Invoke(client *sdk.Client) (response *AddLiveDetectNotifyConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddLiveDetectNotifyConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "AddLiveDetectNotifyConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		AddLiveDetectNotifyConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddLiveDetectNotifyConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddLiveDetectNotifyConfigResponse struct {
	RequestId string
}
