package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	Code      common.String
	Message   common.String
	Data      common.String
	RequestId common.String
	Success   common.String
}
