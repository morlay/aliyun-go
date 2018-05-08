package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteLiveDomainMappingRequest struct {
	requests.RpcRequest
	PullDomain    string `position:"Query" name:"PullDomain"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	PushDomain    string `position:"Query" name:"PushDomain"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (req *DeleteLiveDomainMappingRequest) Invoke(client *sdk.Client) (resp *DeleteLiveDomainMappingResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DeleteLiveDomainMapping", "", "")
	resp = &DeleteLiveDomainMappingResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteLiveDomainMappingResponse struct {
	responses.BaseResponse
	RequestId common.String
}
