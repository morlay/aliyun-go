package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateLiveDetectNotifyConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	NotifyUrl     string `position:"Query" name:"NotifyUrl"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r UpdateLiveDetectNotifyConfigRequest) Invoke(client *sdk.Client) (response *UpdateLiveDetectNotifyConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UpdateLiveDetectNotifyConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "UpdateLiveDetectNotifyConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		UpdateLiveDetectNotifyConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UpdateLiveDetectNotifyConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type UpdateLiveDetectNotifyConfigResponse struct {
	RequestId string
}
