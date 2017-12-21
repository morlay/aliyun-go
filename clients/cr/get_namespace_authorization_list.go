package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetNamespaceAuthorizationListRequest struct {
	requests.RoaRequest
	Namespace string `position:"Path" name:"Namespace"`
}

func (req *GetNamespaceAuthorizationListRequest) Invoke(client *sdk.Client) (resp *GetNamespaceAuthorizationListResponse, err error) {
	req.InitWithApiInfo("cr", "2016-06-07", "GetNamespaceAuthorizationList", "/namespace/[Namespace]/authorizations", "", "")
	req.Method = "GET"

	resp = &GetNamespaceAuthorizationListResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetNamespaceAuthorizationListResponse struct {
	responses.BaseResponse
}
