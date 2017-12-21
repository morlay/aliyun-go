package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteCustomMetricRequest struct {
	GroupId    string `position:"Query" name:"GroupId"`
	MetricName string `position:"Query" name:"MetricName"`
	UUID       string `position:"Query" name:"UUID"`
	Md5        string `position:"Query" name:"Md.5"`
}

func (r DeleteCustomMetricRequest) Invoke(client *sdk.Client) (response *DeleteCustomMetricResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteCustomMetricRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cms", "2017-03-01", "DeleteCustomMetric", "cms", "")

	resp := struct {
		*responses.BaseResponse
		DeleteCustomMetricResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteCustomMetricResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteCustomMetricResponse struct {
	Code      string
	Message   string
	RequestId string
	Result    string
}
