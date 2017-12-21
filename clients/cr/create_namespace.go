package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateNamespaceRequest struct {
	requests.RoaRequest
}

func (req *CreateNamespaceRequest) Invoke(client *sdk.Client) (resp *CreateNamespaceResponse, err error) {
	req.InitWithApiInfo("cr", "2016-06-07", "CreateNamespace", "/namespace", "", "")
	req.Method = "PUT"

	resp = &CreateNamespaceResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateNamespaceResponse struct {
	responses.BaseResponse
}
