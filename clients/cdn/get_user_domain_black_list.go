package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetUserDomainBlackListRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerAccount  string `position:"Query" name:"OwnerAccount"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *GetUserDomainBlackListRequest) Invoke(client *sdk.Client) (resp *GetUserDomainBlackListResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "GetUserDomainBlackList", "", "")
	resp = &GetUserDomainBlackListResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetUserDomainBlackListResponse struct {
	responses.BaseResponse
	RequestId common.String
	Bind      common.String
}
