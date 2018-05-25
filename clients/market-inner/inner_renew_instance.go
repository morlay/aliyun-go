package market_inner

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type InnerRenewInstanceRequest struct {
	requests.RpcRequest
	PayLoad     string `position:"Query" name:"PayLoad"`
	ClientToken string `position:"Query" name:"ClientToken"`
}

func (req *InnerRenewInstanceRequest) Invoke(client *sdk.Client) (resp *InnerRenewInstanceResponse, err error) {
	req.InitWithApiInfo("Market-Inner", "2016-08-01", "InnerRenewInstance", "yunmarket", "")
	resp = &InnerRenewInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type InnerRenewInstanceResponse struct {
	responses.BaseResponse
	RequestId common.String
}
