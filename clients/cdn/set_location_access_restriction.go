package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetLocationAccessRestrictionRequest struct {
	requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	Location      string `position:"Query" name:"Location"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Type          string `position:"Query" name:"Type"`
}

func (req *SetLocationAccessRestrictionRequest) Invoke(client *sdk.Client) (resp *SetLocationAccessRestrictionResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetLocationAccessRestriction", "", "")
	resp = &SetLocationAccessRestrictionResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetLocationAccessRestrictionResponse struct {
	responses.BaseResponse
	RequestId string
}
