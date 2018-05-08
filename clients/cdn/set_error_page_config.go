package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SetErrorPageConfigRequest struct {
	requests.RpcRequest
	PageType      string `position:"Query" name:"PageType"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	CustomPageUrl string `position:"Query" name:"CustomPageUrl"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *SetErrorPageConfigRequest) Invoke(client *sdk.Client) (resp *SetErrorPageConfigResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetErrorPageConfig", "", "")
	resp = &SetErrorPageConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetErrorPageConfigResponse struct {
	responses.BaseResponse
	RequestId common.String
}
