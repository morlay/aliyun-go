package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddCdnDomainRequest struct {
	TopLevelDomain  string `position:"Query" name:"TopLevelDomain"`
	Sources         string `position:"Query" name:"Sources"`
	OwnerAccount    string `position:"Query" name:"OwnerAccount"`
	DomainName      string `position:"Query" name:"DomainName"`
	LiveType        string `position:"Query" name:"LiveType"`
	OwnerId         int64  `position:"Query" name:"OwnerId"`
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	SourcePort      int    `position:"Query" name:"SourcePort"`
	Priorities      string `position:"Query" name:"Priorities"`
	SecurityToken   string `position:"Query" name:"SecurityToken"`
	CdnType         string `position:"Query" name:"CdnType"`
	Scope           string `position:"Query" name:"Scope"`
	SourceType      string `position:"Query" name:"SourceType"`
	CheckUrl        string `position:"Query" name:"CheckUrl"`
	Region          string `position:"Query" name:"Region"`
}

func (r AddCdnDomainRequest) Invoke(client *sdk.Client) (response *AddCdnDomainResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddCdnDomainRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "AddCdnDomain", "", "")

	resp := struct {
		*responses.BaseResponse
		AddCdnDomainResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddCdnDomainResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddCdnDomainResponse struct {
	RequestId string
}
