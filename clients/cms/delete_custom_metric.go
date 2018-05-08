package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteCustomMetricRequest struct {
	requests.RpcRequest
	GroupId    string `position:"Query" name:"GroupId"`
	MetricName string `position:"Query" name:"MetricName"`
	UUID       string `position:"Query" name:"UUID"`
	Md5        string `position:"Query" name:"Md.5"`
}

func (req *DeleteCustomMetricRequest) Invoke(client *sdk.Client) (resp *DeleteCustomMetricResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "DeleteCustomMetric", "cms", "")
	resp = &DeleteCustomMetricResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteCustomMetricResponse struct {
	responses.BaseResponse
	Code      common.String
	Message   common.String
	RequestId common.String
	Result    common.String
}
