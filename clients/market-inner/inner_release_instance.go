package market_inner

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type InnerReleaseInstanceRequest struct {
	requests.RpcRequest
	PayLoad     string `position:"Query" name:"PayLoad"`
	ClientToken string `position:"Query" name:"ClientToken"`
}

func (req *InnerReleaseInstanceRequest) Invoke(client *sdk.Client) (resp *InnerReleaseInstanceResponse, err error) {
	req.InitWithApiInfo("Market-Inner", "2016-08-01", "InnerReleaseInstance", "yunmarket", "")
	resp = &InnerReleaseInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type InnerReleaseInstanceResponse struct {
	responses.BaseResponse
	RequestId common.String
}
