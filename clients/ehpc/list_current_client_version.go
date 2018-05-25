package ehpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListCurrentClientVersionRequest struct {
	requests.RpcRequest
}

func (req *ListCurrentClientVersionRequest) Invoke(client *sdk.Client) (resp *ListCurrentClientVersionResponse, err error) {
	req.InitWithApiInfo("EHPC", "2018-04-12", "ListCurrentClientVersion", "ehs", "")
	resp = &ListCurrentClientVersionResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListCurrentClientVersionResponse struct {
	responses.BaseResponse
	RequestId     common.String
	ClientVersion common.String
}
