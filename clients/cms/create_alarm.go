package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CreateAlarmRequest struct {
	Callby_cms_owner   string `position:"Query" name:"Callby_cms_owner"`
	Period             int    `position:"Query" name:"Period"`
	Webhook            string `position:"Query" name:"Webhook"`
	ContactGroups      string `position:"Query" name:"ContactGroups"`
	EndTime            int    `position:"Query" name:"EndTime"`
	Threshold          string `position:"Query" name:"Threshold"`
	StartTime          int    `position:"Query" name:"StartTime"`
	Name               string `position:"Query" name:"Name"`
	Namespace          string `position:"Query" name:"Namespace"`
	EvaluationCount    int    `position:"Query" name:"EvaluationCount"`
	SilenceTime        int    `position:"Query" name:"SilenceTime"`
	MetricName         string `position:"Query" name:"MetricName"`
	NotifyType         int    `position:"Query" name:"NotifyType"`
	ComparisonOperator string `position:"Query" name:"ComparisonOperator"`
	Dimensions         string `position:"Query" name:"Dimensions"`
	Statistics         string `position:"Query" name:"Statistics"`
}

func (r CreateAlarmRequest) Invoke(client *sdk.Client) (response *CreateAlarmResponse, err error) {
	req := struct {
		*requests.RpcRequest
		CreateAlarmRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cms", "2017-03-01", "CreateAlarm", "cms", "")

	resp := struct {
		*responses.BaseResponse
		CreateAlarmResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.CreateAlarmResponse

	err = client.DoAction(&req, &resp)
	return
}

type CreateAlarmResponse struct {
	Success   bool
	Code      string
	Message   string
	RequestId string
	Data      string
}
