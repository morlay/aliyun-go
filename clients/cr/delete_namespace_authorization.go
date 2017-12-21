package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteNamespaceAuthorizationRequest struct {
	Namespace   string `position:"Path" name:"Namespace"`
	AuthorizeId int64  `position:"Path" name:"AuthorizeId"`
}

func (r DeleteNamespaceAuthorizationRequest) Invoke(client *sdk.Client) (response *DeleteNamespaceAuthorizationResponse, err error) {
	req := struct {
		*requests.RoaRequest
		DeleteNamespaceAuthorizationRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("cr", "2016-06-07", "DeleteNamespaceAuthorization", "/namespace/[Namespace]/authorizations/[AuthorizeId]", "", "")
	req.Method = "DELETE"

	resp := struct {
		*responses.BaseResponse
		DeleteNamespaceAuthorizationResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteNamespaceAuthorizationResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteNamespaceAuthorizationResponse struct {
}
