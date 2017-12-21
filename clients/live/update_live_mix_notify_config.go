package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateLiveMixNotifyConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	NotifyUrl     string `position:"Query" name:"NotifyUrl"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r UpdateLiveMixNotifyConfigRequest) Invoke(client *sdk.Client) (response *UpdateLiveMixNotifyConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UpdateLiveMixNotifyConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "UpdateLiveMixNotifyConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		UpdateLiveMixNotifyConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UpdateLiveMixNotifyConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type UpdateLiveMixNotifyConfigResponse struct {
	RequestId string
}
