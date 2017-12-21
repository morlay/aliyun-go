package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetRefererConfigRequest struct {
	requests.RpcRequest
	ReferList     string `position:"Query" name:"ReferList"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	ReferType     string `position:"Query" name:"ReferType"`
	DisableAst    string `position:"Query" name:"DisableAst"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	AllowEmpty    string `position:"Query" name:"AllowEmpty"`
}

func (req *SetRefererConfigRequest) Invoke(client *sdk.Client) (resp *SetRefererConfigResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetRefererConfig", "", "")
	resp = &SetRefererConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetRefererConfigResponse struct {
	responses.BaseResponse
	RequestId string
}
