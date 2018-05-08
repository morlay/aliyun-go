package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryCustomEventCountRequest struct {
	requests.RpcRequest
	QueryJson string `position:"Query" name:"QueryJson"`
}

func (req *QueryCustomEventCountRequest) Invoke(client *sdk.Client) (resp *QueryCustomEventCountResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "QueryCustomEventCount", "cms", "")
	resp = &QueryCustomEventCountResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryCustomEventCountResponse struct {
	responses.BaseResponse
	Code      common.String
	Message   common.String
	Data      common.String
	RequestId common.String
	Success   bool
}
