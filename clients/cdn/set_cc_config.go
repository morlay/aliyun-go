package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SetCcConfigRequest struct {
	requests.RpcRequest
	AllowIps      string `position:"Query" name:"AllowIps"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	BlockIps      string `position:"Query" name:"BlockIps"`
}

func (req *SetCcConfigRequest) Invoke(client *sdk.Client) (resp *SetCcConfigResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetCcConfig", "", "")
	resp = &SetCcConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetCcConfigResponse struct {
	responses.BaseResponse
	RequestId common.String
}
