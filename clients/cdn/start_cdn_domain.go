package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type StartCdnDomainRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r StartCdnDomainRequest) Invoke(client *sdk.Client) (response *StartCdnDomainResponse, err error) {
	req := struct {
		*requests.RpcRequest
		StartCdnDomainRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "StartCdnDomain", "", "")

	resp := struct {
		*responses.BaseResponse
		StartCdnDomainResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.StartCdnDomainResponse

	err = client.DoAction(&req, &resp)
	return
}

type StartCdnDomainResponse struct {
	RequestId string
}
