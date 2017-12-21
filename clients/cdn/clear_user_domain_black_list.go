package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ClearUserDomainBlackListRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerAccount  string `position:"Query" name:"OwnerAccount"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *ClearUserDomainBlackListRequest) Invoke(client *sdk.Client) (resp *ClearUserDomainBlackListResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "ClearUserDomainBlackList", "", "")
	resp = &ClearUserDomainBlackListResponse{}
	err = client.DoAction(req, resp)
	return
}

type ClearUserDomainBlackListResponse struct {
	responses.BaseResponse
	RequestId string
}
