package ehpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListCurrentClientVersionRequest struct {
	requests.RpcRequest
}

func (req *ListCurrentClientVersionRequest) Invoke(client *sdk.Client) (resp *ListCurrentClientVersionResponse, err error) {
	req.InitWithApiInfo("EHPC", "2017-07-14", "ListCurrentClientVersion", "ehs", "")
	resp = &ListCurrentClientVersionResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListCurrentClientVersionResponse struct {
	responses.BaseResponse
	RequestId     string
	ClientVersion string
}
