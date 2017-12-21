package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetUserDomainBlackListRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerAccount  string `position:"Query" name:"OwnerAccount"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *SetUserDomainBlackListRequest) Invoke(client *sdk.Client) (resp *SetUserDomainBlackListResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetUserDomainBlackList", "", "")
	resp = &SetUserDomainBlackListResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetUserDomainBlackListResponse struct {
	responses.BaseResponse
	RequestId string
}
