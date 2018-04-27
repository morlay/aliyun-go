package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryCustomEventHistogramRequest struct {
	requests.RpcRequest
	QueryJson string `position:"Query" name:"QueryJson"`
}

func (req *QueryCustomEventHistogramRequest) Invoke(client *sdk.Client) (resp *QueryCustomEventHistogramResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "QueryCustomEventHistogram", "cms", "")
	resp = &QueryCustomEventHistogramResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryCustomEventHistogramResponse struct {
	responses.BaseResponse
	Code      string
	Message   string
	Data      string
	RequestId string
	Success   string
}
