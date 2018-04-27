package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	Code      string
	Message   string
	Data      string
	RequestId string
	Success   bool
}
