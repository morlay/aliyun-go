package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SetIpAllowListConfigRequest struct {
	requests.RpcRequest
	AllowIps      string `position:"Query" name:"AllowIps"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *SetIpAllowListConfigRequest) Invoke(client *sdk.Client) (resp *SetIpAllowListConfigResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetIpAllowListConfig", "", "")
	resp = &SetIpAllowListConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetIpAllowListConfigResponse struct {
	responses.BaseResponse
	RequestId common.String
}
