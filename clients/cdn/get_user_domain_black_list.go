package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetUserDomainBlackListRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerAccount  string `position:"Query" name:"OwnerAccount"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r GetUserDomainBlackListRequest) Invoke(client *sdk.Client) (response *GetUserDomainBlackListResponse, err error) {
	req := struct {
		*requests.RpcRequest
		GetUserDomainBlackListRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "GetUserDomainBlackList", "", "")

	resp := struct {
		*responses.BaseResponse
		GetUserDomainBlackListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetUserDomainBlackListResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetUserDomainBlackListResponse struct {
	RequestId string
	Bind      string
}
