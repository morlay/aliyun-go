package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyCdnDomainRequest struct {
	requests.RpcRequest
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

func (req *ModifyCdnDomainRequest) Invoke(client *sdk.Client) (resp *ModifyCdnDomainResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "ModifyCdnDomain", "", "")
	resp = &ModifyCdnDomainResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyCdnDomainResponse struct {
	responses.BaseResponse
	RequestId string
}
