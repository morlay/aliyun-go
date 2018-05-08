package dcdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type StartDcdnDomainRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *StartDcdnDomainRequest) Invoke(client *sdk.Client) (resp *StartDcdnDomainResponse, err error) {
	req.InitWithApiInfo("dcdn", "2018-01-15", "StartDcdnDomain", "dcdn", "")
	resp = &StartDcdnDomainResponse{}
	err = client.DoAction(req, resp)
	return
}

type StartDcdnDomainResponse struct {
	responses.BaseResponse
	RequestId common.String
}
