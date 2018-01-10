package csb

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateServiceListStatusRequest struct {
	requests.RpcRequest
	Data  string `position:"Body" name:"Data"`
	CsbId int64  `position:"Query" name:"CsbId"`
}

func (req *UpdateServiceListStatusRequest) Invoke(client *sdk.Client) (resp *UpdateServiceListStatusResponse, err error) {
	req.InitWithApiInfo("CSB", "2017-11-18", "UpdateServiceListStatus", "CSB", "")
	resp = &UpdateServiceListStatusResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateServiceListStatusResponse struct {
	responses.BaseResponse
	Code      int
	Message   string
	RequestId string
}
