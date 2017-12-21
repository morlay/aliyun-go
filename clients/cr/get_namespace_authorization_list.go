package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetNamespaceAuthorizationListRequest struct {
	Namespace string `position:"Path" name:"Namespace"`
}

func (r GetNamespaceAuthorizationListRequest) Invoke(client *sdk.Client) (response *GetNamespaceAuthorizationListResponse, err error) {
	req := struct {
		*requests.RoaRequest
		GetNamespaceAuthorizationListRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("cr", "2016-06-07", "GetNamespaceAuthorizationList", "/namespace/[Namespace]/authorizations", "", "")
	req.Method = "GET"

	resp := struct {
		*responses.BaseResponse
		GetNamespaceAuthorizationListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.GetNamespaceAuthorizationListResponse

	err = client.DoAction(&req, &resp)
	return
}

type GetNamespaceAuthorizationListResponse struct {
}
