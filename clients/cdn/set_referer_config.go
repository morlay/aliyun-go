package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetRefererConfigRequest struct {
	ReferList     string `position:"Query" name:"ReferList"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	ReferType     string `position:"Query" name:"ReferType"`
	DisableAst    string `position:"Query" name:"DisableAst"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	AllowEmpty    string `position:"Query" name:"AllowEmpty"`
}

func (r SetRefererConfigRequest) Invoke(client *sdk.Client) (response *SetRefererConfigResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetRefererConfigRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetRefererConfig", "", "")

	resp := struct {
		*responses.BaseResponse
		SetRefererConfigResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetRefererConfigResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetRefererConfigResponse struct {
	RequestId string
}
