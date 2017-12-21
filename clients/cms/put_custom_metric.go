package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type PutCustomMetricRequest struct {
	MetricList string `position:"Query" name:"MetricList"`
}

func (r PutCustomMetricRequest) Invoke(client *sdk.Client) (response *PutCustomMetricResponse, err error) {
	req := struct {
		*requests.RpcRequest
		PutCustomMetricRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cms", "2017-03-01", "PutCustomMetric", "cms", "")

	resp := struct {
		*responses.BaseResponse
		PutCustomMetricResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.PutCustomMetricResponse

	err = client.DoAction(&req, &resp)
	return
}

type PutCustomMetricResponse struct {
	Code    string
	Message string
	Data    string
}
