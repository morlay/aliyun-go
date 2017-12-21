package cms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type UpdateAlarmRequest struct {
	Callby_cms_owner   string `position:"Query" name:"Callby_cms_owner"`
	Period             int    `position:"Query" name:"Period"`
	Webhook            string `position:"Query" name:"Webhook"`
	ContactGroups      string `position:"Query" name:"ContactGroups"`
	EndTime            int    `position:"Query" name:"EndTime"`
	Threshold          string `position:"Query" name:"Threshold"`
	StartTime          int    `position:"Query" name:"StartTime"`
	Name               string `position:"Query" name:"Name"`
	EvaluationCount    int    `position:"Query" name:"EvaluationCount"`
	SilenceTime        int    `position:"Query" name:"SilenceTime"`
	Id                 string `position:"Query" name:"Id"`
	NotifyType         int    `position:"Query" name:"NotifyType"`
	ComparisonOperator string `position:"Query" name:"ComparisonOperator"`
	Statistics         string `position:"Query" name:"Statistics"`
}

func (r UpdateAlarmRequest) Invoke(client *sdk.Client) (response *UpdateAlarmResponse, err error) {
	req := struct {
		*requests.RpcRequest
		UpdateAlarmRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cms", "2017-03-01", "UpdateAlarm", "cms", "")

	resp := struct {
		*responses.BaseResponse
		UpdateAlarmResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.UpdateAlarmResponse

	err = client.DoAction(&req, &resp)
	return
}

type UpdateAlarmResponse struct {
	Success   bool
	Code      string
	Message   string
	RequestId string
}
