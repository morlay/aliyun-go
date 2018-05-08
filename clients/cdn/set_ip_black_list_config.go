package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SetIpBlackListConfigRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	BlockIps      string `position:"Query" name:"BlockIps"`
}

func (req *SetIpBlackListConfigRequest) Invoke(client *sdk.Client) (resp *SetIpBlackListConfigResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetIpBlackListConfig", "", "")
	resp = &SetIpBlackListConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetIpBlackListConfigResponse struct {
	responses.BaseResponse
	RequestId common.String
}
