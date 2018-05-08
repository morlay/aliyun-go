package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyHttpHeaderConfigRequest struct {
	requests.RpcRequest
	HeaderValue   string `position:"Query" name:"HeaderValue"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	ConfigID      string `position:"Query" name:"ConfigID"`
	DomainName    string `position:"Query" name:"DomainName"`
	HeaderKey     string `position:"Query" name:"HeaderKey"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifyHttpHeaderConfigRequest) Invoke(client *sdk.Client) (resp *ModifyHttpHeaderConfigResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "ModifyHttpHeaderConfig", "", "")
	resp = &ModifyHttpHeaderConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyHttpHeaderConfigResponse struct {
	responses.BaseResponse
	RequestId common.String
}
