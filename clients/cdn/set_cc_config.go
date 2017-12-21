package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetCcConfigRequest struct {
	AllowIps      string `position:"Query" name:"AllowIps"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	BlockIps      string `position:"Query" name:"BlockIps"`
}

func (r SetCcConfigRequest) Invoke(client *sdk.Client) (response *SetCcConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetCcConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetCcConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		SetCcConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetCcConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetCcConfigResponse struct {
	RequestId string
}
