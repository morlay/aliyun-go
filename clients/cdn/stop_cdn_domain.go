package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type StopCdnDomainRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r StopCdnDomainRequest) Invoke(client *sdk.Client) (response *StopCdnDomainResponse, err error) {
	req := struct {
		*requests.RpcRequest
		StopCdnDomainRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "StopCdnDomain", "", "")

	resp := struct {
		*responses.BaseResponse
		StopCdnDomainResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.StopCdnDomainResponse

	err = client.DoAction(&req, &resp)
	return
}

type StopCdnDomainResponse struct {
	RequestId string
}
