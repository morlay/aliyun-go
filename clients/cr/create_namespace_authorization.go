package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateNamespaceAuthorizationRequest struct {
	Namespace string `position:"Path" name:"Namespace"`
}

func (r CreateNamespaceAuthorizationRequest) Invoke(client *sdk.Client) (response *CreateNamespaceAuthorizationResponse, err error) {
	req := struct {
		*requests.RoaRequest
		CreateNamespaceAuthorizationRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("cr", "2016-06-07", "CreateNamespaceAuthorization", "/namespace/[Namespace]/authorizations", "", "")
	req.Method = "PUT"

	resp := struct {
		*responses.BaseResponse
		CreateNamespaceAuthorizationResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateNamespaceAuthorizationResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateNamespaceAuthorizationResponse struct {
}
