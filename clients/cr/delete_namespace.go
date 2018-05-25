package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteNamespaceRequest struct {
	requests.RoaRequest
	Namespace string `position:"Path" name:"Namespace"`
}

func (req *DeleteNamespaceRequest) Invoke(client *sdk.Client) (resp *DeleteNamespaceResponse, err error) {
	req.InitWithApiInfo("cr", "2016-06-07", "DeleteNamespace", "/namespace/[Namespace]", "", "")
	req.Method = "DELETE"

	resp = &DeleteNamespaceResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteNamespaceResponse struct {
	responses.BaseResponse
}
