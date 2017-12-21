package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetWafConfigRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	Enable        string `position:"Query" name:"Enable"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r SetWafConfigRequest) Invoke(client *sdk.Client) (response *SetWafConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetWafConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetWafConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		SetWafConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetWafConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetWafConfigResponse struct {
	RequestId string
}
