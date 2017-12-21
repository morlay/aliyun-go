package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteLiveDomainMappingRequest struct {
	PullDomain    string `position:"Query" name:"PullDomain"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	PushDomain    string `position:"Query" name:"PushDomain"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DeleteLiveDomainMappingRequest) Invoke(client *sdk.Client) (response *DeleteLiveDomainMappingResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteLiveDomainMappingRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DeleteLiveDomainMapping", "", "")

	resp := struct {
		*responses.BaseResponse
		DeleteLiveDomainMappingResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteLiveDomainMappingResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteLiveDomainMappingResponse struct {
	RequestId string
}
