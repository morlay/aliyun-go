package market_inner

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type InnerUpdateInstanceRequest struct {
	requests.RpcRequest
	PayLoad     string `position:"Query" name:"PayLoad"`
	ClientToken string `position:"Query" name:"ClientToken"`
}

func (req *InnerUpdateInstanceRequest) Invoke(client *sdk.Client) (resp *InnerUpdateInstanceResponse, err error) {
	req.InitWithApiInfo("Market-Inner", "2016-08-01", "InnerUpdateInstance", "yunmarket", "")
	resp = &InnerUpdateInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type InnerUpdateInstanceResponse struct {
	responses.BaseResponse
	RequestId common.String
}
