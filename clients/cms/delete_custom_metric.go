package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteCustomMetricRequest struct {
	requests.RpcRequest
	GroupId    string `position:"Query" name:"GroupId"`
	MetricName string `position:"Query" name:"MetricName"`
	UUID       string `position:"Query" name:"UUID"`
	Md5        string `position:"Query" name:"Md.5"`
}

func (req *DeleteCustomMetricRequest) Invoke(client *sdk.Client) (resp *DeleteCustomMetricResponse, err error) {
	req.InitWithApiInfo("Cms", "2017-03-01", "DeleteCustomMetric", "cms", "")
	resp = &DeleteCustomMetricResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteCustomMetricResponse struct {
	responses.BaseResponse
	Code      string
	Message   string
	RequestId string
	Result    string
}
