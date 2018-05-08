package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyDomainCustomLogConfigRequest struct {
	requests.RpcRequest
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Version       string `position:"Query" name:"Version"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	ConfigId      string `position:"Query" name:"ConfigId"`
}

func (req *ModifyDomainCustomLogConfigRequest) Invoke(client *sdk.Client) (resp *ModifyDomainCustomLogConfigResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "ModifyDomainCustomLogConfig", "", "")
	resp = &ModifyDomainCustomLogConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyDomainCustomLogConfigResponse struct {
	responses.BaseResponse
	RequestId common.String
}
