package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateNamespaceRequest struct {
	requests.RoaRequest
	Namespace string `position:"Path" name:"Namespace"`
}

func (req *UpdateNamespaceRequest) Invoke(client *sdk.Client) (resp *UpdateNamespaceResponse, err error) {
	req.InitWithApiInfo("cr", "2016-06-07", "UpdateNamespace", "/namespace/[Namespace]", "", "")
	req.Method = "POST"

	resp = &UpdateNamespaceResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateNamespaceResponse struct {
	responses.BaseResponse
}
