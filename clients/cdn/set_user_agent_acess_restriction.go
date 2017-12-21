package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetUserAgentAcessRestrictionRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	UserAgent     string `position:"Query" name:"UserAgent"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Type          string `position:"Query" name:"Type"`
}

func (req *SetUserAgentAcessRestrictionRequest) Invoke(client *sdk.Client) (resp *SetUserAgentAcessRestrictionResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetUserAgentAcessRestriction", "", "")
	resp = &SetUserAgentAcessRestrictionResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetUserAgentAcessRestrictionResponse struct {
	responses.BaseResponse
	RequestId string
}
