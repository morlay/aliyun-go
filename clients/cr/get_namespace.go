package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetNamespaceRequest struct {
	requests.RoaRequest
	Namespace string `position:"Path" name:"Namespace"`
}

func (req *GetNamespaceRequest) Invoke(client *sdk.Client) (resp *GetNamespaceResponse, err error) {
	req.InitWithApiInfo("cr", "2016-06-07", "GetNamespace", "/namespace/[Namespace]", "", "")
	req.Method = "GET"

	resp = &GetNamespaceResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetNamespaceResponse struct {
	responses.BaseResponse
}
