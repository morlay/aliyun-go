package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddLiveRecordNotifyConfigRequest struct {
	SecurityToken    string `position:"Query" name:"SecurityToken"`
	DomainName       string `position:"Query" name:"DomainName"`
	NotifyUrl        string `position:"Query" name:"NotifyUrl"`
	OwnerId          int64  `position:"Query" name:"OwnerId"`
	NeedStatusNotify string `position:"Query" name:"NeedStatusNotify"`
}

func (r AddLiveRecordNotifyConfigRequest) Invoke(client *sdk.Client) (response *AddLiveRecordNotifyConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddLiveRecordNotifyConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "AddLiveRecordNotifyConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		AddLiveRecordNotifyConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddLiveRecordNotifyConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddLiveRecordNotifyConfigResponse struct {
	RequestId string
}
