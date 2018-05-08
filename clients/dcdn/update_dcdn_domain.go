package dcdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateDcdnDomainRequest struct {
	requests.RpcRequest
	TopLevelDomain  string `position:"Query" name:"TopLevelDomain"`
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	Sources         string `position:"Query" name:"Sources"`
	SecurityToken   string `position:"Query" name:"SecurityToken"`
	DomainName      string `position:"Query" name:"DomainName"`
	OwnerId         int64  `position:"Query" name:"OwnerId"`
}

func (req *UpdateDcdnDomainRequest) Invoke(client *sdk.Client) (resp *UpdateDcdnDomainResponse, err error) {
	req.InitWithApiInfo("dcdn", "2018-01-15", "UpdateDcdnDomain", "dcdn", "")
	resp = &UpdateDcdnDomainResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateDcdnDomainResponse struct {
	responses.BaseResponse
	RequestId string
}
