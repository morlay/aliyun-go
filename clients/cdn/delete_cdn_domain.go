package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteCdnDomainRequest struct {
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	SecurityToken   string `position:"Query" name:"SecurityToken"`
	DomainName      string `position:"Query" name:"DomainName"`
	OwnerId         int64  `position:"Query" name:"OwnerId"`
}

func (r DeleteCdnDomainRequest) Invoke(client *sdk.Client) (response *DeleteCdnDomainResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteCdnDomainRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DeleteCdnDomain", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteCdnDomainResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteCdnDomainResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteCdnDomainResponse struct {
	RequestId string
}
