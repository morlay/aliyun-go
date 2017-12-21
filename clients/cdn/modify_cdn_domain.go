package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyCdnDomainRequest struct {
	TopLevelDomain  string `position:"Query" name:"TopLevelDomain"`
	SourcePort      int    `position:"Query" name:"SourcePort"`
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	Priorities      string `position:"Query" name:"Priorities"`
	Sources         string `position:"Query" name:"Sources"`
	SecurityToken   string `position:"Query" name:"SecurityToken"`
	DomainName      string `position:"Query" name:"DomainName"`
	SourceType      string `position:"Query" name:"SourceType"`
	OwnerId         int64  `position:"Query" name:"OwnerId"`
}

func (r ModifyCdnDomainRequest) Invoke(client *sdk.Client) (response *ModifyCdnDomainResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifyCdnDomainRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "ModifyCdnDomain", "", "")

	resp := struct {
		*responses.BaseResponse
		ModifyCdnDomainResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifyCdnDomainResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifyCdnDomainResponse struct {
	RequestId string
}
