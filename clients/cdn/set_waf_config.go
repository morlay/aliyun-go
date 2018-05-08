package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SetWafConfigRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	Enable        string `position:"Query" name:"Enable"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *SetWafConfigRequest) Invoke(client *sdk.Client) (resp *SetWafConfigResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetWafConfig", "", "")
	resp = &SetWafConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetWafConfigResponse struct {
	responses.BaseResponse
	RequestId common.String
}
