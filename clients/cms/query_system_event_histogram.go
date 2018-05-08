package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QuerySystemEventHistogramRequest struct {
	requests.RpcRequest
	QueryJson string `position:"Query" name:"QueryJson"`
}

func (req *QuerySystemEventHistogramRequest) Invoke(client *sdk.Client) (resp *QuerySystemEventHistogramResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "QuerySystemEventHistogram", "cms", "")
	resp = &QuerySystemEventHistogramResponse{}
	err = client.DoAction(req, resp)
	return
}

type QuerySystemEventHistogramResponse struct {
	responses.BaseResponse
	Code      common.String
	Message   common.String
	Data      common.String
	RequestId common.String
	Success   common.String
}
