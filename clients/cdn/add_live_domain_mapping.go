package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddLiveDomainMappingRequest struct {
	PullDomain    string `position:"Query" name:"PullDomain"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	PushDomain    string `position:"Query" name:"PushDomain"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r AddLiveDomainMappingRequest) Invoke(client *sdk.Client) (response *AddLiveDomainMappingResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddLiveDomainMappingRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "AddLiveDomainMapping", "", "")

	resp := struct {
		*responses.BaseResponse
		AddLiveDomainMappingResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddLiveDomainMappingResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddLiveDomainMappingResponse struct {
	RequestId string
}
