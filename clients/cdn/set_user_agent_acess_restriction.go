package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetUserAgentAcessRestrictionRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	UserAgent     string `position:"Query" name:"UserAgent"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Type          string `position:"Query" name:"Type"`
}

func (r SetUserAgentAcessRestrictionRequest) Invoke(client *sdk.Client) (response *SetUserAgentAcessRestrictionResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetUserAgentAcessRestrictionRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetUserAgentAcessRestriction", "", "")

	resp := struct {
		*responses.BaseResponse
		SetUserAgentAcessRestrictionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetUserAgentAcessRestrictionResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetUserAgentAcessRestrictionResponse struct {
	RequestId string
}
