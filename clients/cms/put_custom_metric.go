package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type PutCustomMetricRequest struct {
	requests.RpcRequest
	MetricList string `position:"Query" name:"MetricList"`
}

func (req *PutCustomMetricRequest) Invoke(client *sdk.Client) (resp *PutCustomMetricResponse, err error) {
	req.InitWithApiInfo("Cms", "2017-03-01", "PutCustomMetric", "cms", "")
	resp = &PutCustomMetricResponse{}
	err = client.DoAction(req, resp)
	return
}

type PutCustomMetricResponse struct {
	responses.BaseResponse
	Code    string
	Message string
	Data    string
}
