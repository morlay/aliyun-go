package csb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ApproveOrderListRequest struct {
	requests.RpcRequest
	Data string `position:"Body" name:"Data"`
}

func (req *ApproveOrderListRequest) Invoke(client *sdk.Client) (resp *ApproveOrderListResponse, err error) {
	req.InitWithApiInfo("CSB", "2017-11-18", "ApproveOrderList", "CSB", "")
	resp = &ApproveOrderListResponse{}
	err = client.DoAction(req, resp)
	return
}

type ApproveOrderListResponse struct {
	responses.BaseResponse
	Code      common.Integer
	Message   common.String
	RequestId common.String
}
