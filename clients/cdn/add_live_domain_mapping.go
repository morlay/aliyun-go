package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddLiveDomainMappingRequest struct {
	requests.RpcRequest
	PullDomain    string `position:"Query" name:"PullDomain"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	PushDomain    string `position:"Query" name:"PushDomain"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *AddLiveDomainMappingRequest) Invoke(client *sdk.Client) (resp *AddLiveDomainMappingResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "AddLiveDomainMapping", "", "")
	resp = &AddLiveDomainMappingResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddLiveDomainMappingResponse struct {
	responses.BaseResponse
	RequestId string
}
