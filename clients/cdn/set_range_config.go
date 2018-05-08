package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SetRangeConfigRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	Enable        string `position:"Query" name:"Enable"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *SetRangeConfigRequest) Invoke(client *sdk.Client) (resp *SetRangeConfigResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetRangeConfig", "", "")
	resp = &SetRangeConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetRangeConfigResponse struct {
	responses.BaseResponse
	RequestId common.String
}
