package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateLiveRecordNotifyConfigRequest struct {
	SecurityToken    string `position:"Query" name:"SecurityToken"`
	DomainName       string `position:"Query" name:"DomainName"`
	NotifyUrl        string `position:"Query" name:"NotifyUrl"`
	OwnerId          int64  `position:"Query" name:"OwnerId"`
	NeedStatusNotify string `position:"Query" name:"NeedStatusNotify"`
}

func (r UpdateLiveRecordNotifyConfigRequest) Invoke(client *sdk.Client) (response *UpdateLiveRecordNotifyConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UpdateLiveRecordNotifyConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "UpdateLiveRecordNotifyConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		UpdateLiveRecordNotifyConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UpdateLiveRecordNotifyConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type UpdateLiveRecordNotifyConfigResponse struct {
	RequestId string
}
