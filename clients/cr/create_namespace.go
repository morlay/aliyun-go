package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateNamespaceRequest struct {
}

func (r CreateNamespaceRequest) Invoke(client *sdk.Client) (response *CreateNamespaceResponse, err error) {
	req := struct {
		*requests.RoaRequest
		CreateNamespaceRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("cr", "2016-06-07", "CreateNamespace", "/namespace", "", "")
	req.Method = "PUT"

	resp := struct {
		*responses.BaseResponse
		CreateNamespaceResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateNamespaceResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateNamespaceResponse struct {
}
