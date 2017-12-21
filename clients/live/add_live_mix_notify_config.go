package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddLiveMixNotifyConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	NotifyUrl     string `position:"Query" name:"NotifyUrl"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r AddLiveMixNotifyConfigRequest) Invoke(client *sdk.Client) (response *AddLiveMixNotifyConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddLiveMixNotifyConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "AddLiveMixNotifyConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		AddLiveMixNotifyConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddLiveMixNotifyConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddLiveMixNotifyConfigResponse struct {
	RequestId string
}
