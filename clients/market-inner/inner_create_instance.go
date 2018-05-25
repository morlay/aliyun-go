package market_inner

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type InnerCreateInstanceRequest struct {
	requests.RpcRequest
	PayLoad     string `position:"Query" name:"PayLoad"`
	ClientToken string `position:"Query" name:"ClientToken"`
}

func (req *InnerCreateInstanceRequest) Invoke(client *sdk.Client) (resp *InnerCreateInstanceResponse, err error) {
	req.InitWithApiInfo("Market-Inner", "2016-08-01", "InnerCreateInstance", "yunmarket", "")
	resp = &InnerCreateInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type InnerCreateInstanceResponse struct {
	responses.BaseResponse
	RequestId  common.String
	InstanceId common.String
}
