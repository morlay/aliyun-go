package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetUserDomainBlackListRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerAccount  string `position:"Query" name:"OwnerAccount"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r SetUserDomainBlackListRequest) Invoke(client *sdk.Client) (response *SetUserDomainBlackListResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetUserDomainBlackListRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetUserDomainBlackList", "", "")

	resp := struct {
		*responses.BaseResponse
		SetUserDomainBlackListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetUserDomainBlackListResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetUserDomainBlackListResponse struct {
	RequestId string
}
