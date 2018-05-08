package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteCdnDomainRequest struct {
	requests.RpcRequest
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	SecurityToken   string `position:"Query" name:"SecurityToken"`
	DomainName      string `position:"Query" name:"DomainName"`
	OwnerId         int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteCdnDomainRequest) Invoke(client *sdk.Client) (resp *DeleteCdnDomainResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DeleteCdnDomain", "", "")
	resp = &DeleteCdnDomainResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteCdnDomainResponse struct {
	responses.BaseResponse
	RequestId common.String
}
