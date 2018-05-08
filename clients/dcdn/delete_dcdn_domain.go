package dcdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteDcdnDomainRequest struct {
	requests.RpcRequest
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	SecurityToken   string `position:"Query" name:"SecurityToken"`
	OwnerAccount    string `position:"Query" name:"OwnerAccount"`
	DomainName      string `position:"Query" name:"DomainName"`
	OwnerId         int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteDcdnDomainRequest) Invoke(client *sdk.Client) (resp *DeleteDcdnDomainResponse, err error) {
	req.InitWithApiInfo("dcdn", "2018-01-15", "DeleteDcdnDomain", "dcdn", "")
	resp = &DeleteDcdnDomainResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteDcdnDomainResponse struct {
	responses.BaseResponse
	RequestId string
}
