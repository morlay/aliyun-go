package dcdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type StopDcdnDomainRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *StopDcdnDomainRequest) Invoke(client *sdk.Client) (resp *StopDcdnDomainResponse, err error) {
	req.InitWithApiInfo("dcdn", "2018-01-15", "StopDcdnDomain", "dcdn", "")
	resp = &StopDcdnDomainResponse{}
	err = client.DoAction(req, resp)
	return
}

type StopDcdnDomainResponse struct {
	responses.BaseResponse
	RequestId common.String
}
