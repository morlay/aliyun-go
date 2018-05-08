package dcdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AddDcdnDomainRequest struct {
	requests.RpcRequest
	TopLevelDomain  string `position:"Query" name:"TopLevelDomain"`
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	Sources         string `position:"Query" name:"Sources"`
	SecurityToken   string `position:"Query" name:"SecurityToken"`
	OwnerAccount    string `position:"Query" name:"OwnerAccount"`
	Scope           string `position:"Query" name:"Scope"`
	DomainName      string `position:"Query" name:"DomainName"`
	OwnerId         int64  `position:"Query" name:"OwnerId"`
	CheckUrl        string `position:"Query" name:"CheckUrl"`
}

func (req *AddDcdnDomainRequest) Invoke(client *sdk.Client) (resp *AddDcdnDomainResponse, err error) {
	req.InitWithApiInfo("dcdn", "2018-01-15", "AddDcdnDomain", "dcdn", "")
	resp = &AddDcdnDomainResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddDcdnDomainResponse struct {
	responses.BaseResponse
	RequestId common.String
}
