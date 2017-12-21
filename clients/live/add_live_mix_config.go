package live

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddLiveMixConfigRequest struct {
	Template      string `position:"Query" name:"Template"`
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r AddLiveMixConfigRequest) Invoke(client *sdk.Client) (response *AddLiveMixConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddLiveMixConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "AddLiveMixConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		AddLiveMixConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddLiveMixConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddLiveMixConfigResponse struct {
	RequestId string
}
