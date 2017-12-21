package cdn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetLocationAccessRestrictionRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	Location      string `position:"Query" name:"Location"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	Type          string `position:"Query" name:"Type"`
}

func (r SetLocationAccessRestrictionRequest) Invoke(client *sdk.Client) (response *SetLocationAccessRestrictionResponse, err error) {
	req := struct {
		*requests.RpcRequest
		SetLocationAccessRestrictionRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "SetLocationAccessRestriction", "", "")

	resp := struct {
		*responses.BaseResponse
		SetLocationAccessRestrictionResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.SetLocationAccessRestrictionResponse

	err = client.DoAction(&req, &resp)
	return
}

type SetLocationAccessRestrictionResponse struct {
	RequestId string
}
