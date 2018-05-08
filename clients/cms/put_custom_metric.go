package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type PutCustomMetricRequest struct {
	requests.RpcRequest
	MetricList string `position:"Query" name:"MetricList"`
}

func (req *PutCustomMetricRequest) Invoke(client *sdk.Client) (resp *PutCustomMetricResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "PutCustomMetric", "cms", "")
	resp = &PutCustomMetricResponse{}
	err = client.DoAction(req, resp)
	return
}

type PutCustomMetricResponse struct {
	responses.BaseResponse
	Code    common.String
	Message common.String
	Data    common.String
}
