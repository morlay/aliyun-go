package market_inner

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type InnerExpiredInstanceRequest struct {
	requests.RpcRequest
	PayLoad     string `position:"Query" name:"PayLoad"`
	ClientToken string `position:"Query" name:"ClientToken"`
}

func (req *InnerExpiredInstanceRequest) Invoke(client *sdk.Client) (resp *InnerExpiredInstanceResponse, err error) {
	req.InitWithApiInfo("Market-Inner", "2016-08-01", "InnerExpiredInstance", "yunmarket", "")
	resp = &InnerExpiredInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type InnerExpiredInstanceResponse struct {
	responses.BaseResponse
	RequestId common.String
}
