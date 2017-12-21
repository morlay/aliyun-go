package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateNamespaceAuthorizationRequest struct {
	requests.RoaRequest
	Namespace string `position:"Path" name:"Namespace"`
}

func (req *CreateNamespaceAuthorizationRequest) Invoke(client *sdk.Client) (resp *CreateNamespaceAuthorizationResponse, err error) {
	req.InitWithApiInfo("cr", "2016-06-07", "CreateNamespaceAuthorization", "/namespace/[Namespace]/authorizations", "", "")
	req.Method = "PUT"

	resp = &CreateNamespaceAuthorizationResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateNamespaceAuthorizationResponse struct {
	responses.BaseResponse
}
