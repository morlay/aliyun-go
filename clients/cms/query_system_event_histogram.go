package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	Code      string
	Message   string
	Data      string
	RequestId string
	Success   string
}
