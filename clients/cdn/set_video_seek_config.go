package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SetVideoSeekConfigRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	Enable        string `position:"Query" name:"Enable"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *SetVideoSeekConfigRequest) Invoke(client *sdk.Client) (resp *SetVideoSeekConfigResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetVideoSeekConfig", "", "")
	resp = &SetVideoSeekConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetVideoSeekConfigResponse struct {
	responses.BaseResponse
	RequestId common.String
}
