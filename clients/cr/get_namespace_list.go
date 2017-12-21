package cr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetNamespaceListRequest struct {
	requests.RoaRequest
}

func (req *GetNamespaceListRequest) Invoke(client *sdk.Client) (resp *GetNamespaceListResponse, err error) {
	req.InitWithApiInfo("cr", "2016-06-07", "GetNamespaceList", "/namespace", "", "")
	req.Method = "GET"

	resp = &GetNamespaceListResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetNamespaceListResponse struct {
	responses.BaseResponse
}
