package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CheckUserBalanceRequest struct {
	requests.RpcRequest
}

func (req *CheckUserBalanceRequest) Invoke(client *sdk.Client) (resp *CheckUserBalanceResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "CheckUserBalance", "", "")
	resp = &CheckUserBalanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type CheckUserBalanceResponse struct {
	responses.BaseResponse
	RequestId string
	Balance   string
	Enough    string
}
